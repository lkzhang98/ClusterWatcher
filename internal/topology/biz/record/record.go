package record

import (
	"ClusterWatcher/internal/pkg/errno"
	"ClusterWatcher/internal/pkg/log"
	"ClusterWatcher/internal/pkg/model"
	"ClusterWatcher/internal/topology/store"
	v1 "ClusterWatcher/pkg/api/topology/v1"
	"context"
	"github.com/pkg/errors"
	"github.com/qiniu/qmgo"
	"time"
)

type RecordBiz interface {
	List(context.Context, string, time.Time, time.Time) (*v1.ListRecordResponse, error)
	GetName(context.Context, time.Time, time.Time) (*v1.ListTopologyNameResponse, error)
	GetNs(context.Context, time.Time, time.Time) (*v1.ListTopologyNsResponse, error)
	Layer(context.Context, time.Time, time.Time) (*v1.ListTopologyLayerResponse, error)
}

type recordBiz struct {
	ds store.IStore
}

var _ RecordBiz = (*recordBiz)(nil)

func New(ds store.IStore) *recordBiz {
	return &recordBiz{ds: ds}
}

func (r *recordBiz) List(ctx context.Context, name string, startTime time.Time, endTime time.Time) (*v1.ListRecordResponse, error) {

	record, err := r.ds.Records().List(ctx, name, startTime, endTime)

	if err != nil {
		if errors.Is(err, qmgo.ErrQueryResultTypeInconsistent) {
			return nil, errno.ErrRecordConnectFail
		}
	}
	var resp v1.ListRecordResponse
	log.Debug("[biz] get data from database successfully")

	resp.Records = record

	if err != nil {
		log.Errorf("[biz] Fail to encode object %v", err)
		return nil, errno.ErrRecordConnectFail
	}
	return &resp, nil
}

func (r *recordBiz) GetName(ctx context.Context, startTime time.Time, endTime time.Time) (*v1.ListTopologyNameResponse, error) {
	record, err := r.ds.Records().List(ctx, "pod", startTime, endTime)
	if err != nil {
		if errors.Is(err, qmgo.ErrQueryResultTypeInconsistent) {
			return nil, errno.ErrRecordConnectFail
		}
	}

	var topologyGroup *model.APITopologyGroup = model.NewAPITopologyGroup()
	for key, value := range record {
		for _, table := range value.Tables {
			if table.ID == "kubernetes_labels_" {
				for _, row := range table.Rows {
					if row.ID == "label_app" {
						group_id := row.Entries["value"]
						if _, ok := topologyGroup.SubTopology[group_id]; ok {
							topologyGroup.SubTopology[group_id].Nodes[key] = value
						} else {
							newSubTopology := model.TopologyGroup{Id: group_id, Nodes: make(model.NodeSummaries)}
							newSubTopology.Nodes[key] = value
							topologyGroup.SubTopology[group_id] = newSubTopology
							topologyGroup.GroupList = append(topologyGroup.GroupList, group_id)
						}
					}
				}
			}
		}
	}

	var resp v1.ListTopologyNameResponse
	resp.Topology = *topologyGroup
	log.Debug("[biz] get data from database successfully")

	log.Debugf("get data %+v", resp)
	return &resp, nil
}

func (r *recordBiz) GetNs(ctx context.Context, startTime time.Time, endTime time.Time) (*v1.ListTopologyNsResponse, error) {
	record, err := r.ds.Records().List(ctx, "pod", startTime, endTime)
	if err != nil {
		if errors.Is(err, qmgo.ErrQueryResultTypeInconsistent) {
			return nil, errno.ErrRecordConnectFail
		}
	}

	var topologyGroup *model.APITopologyGroup = model.NewAPITopologyGroup()
	for key, value := range record {
		for _, meta := range value.Metadata {
			if meta.ID == "kubernetes_namespace" {
				group_id := meta.Value
				if _, ok := topologyGroup.SubTopology[group_id]; ok {
					topologyGroup.SubTopology[group_id].Nodes[key] = value
				} else {
					newSubTopology := model.TopologyGroup{Id: group_id, Nodes: make(model.NodeSummaries)}
					newSubTopology.Nodes[key] = value
					topologyGroup.SubTopology[group_id] = newSubTopology
					topologyGroup.GroupList = append(topologyGroup.GroupList, group_id)
				}
				break
			}
		}
	}

	var resp v1.ListTopologyNsResponse
	resp.Topology = *topologyGroup
	log.Debug("[biz] get data from database successfully")

	log.Debugf("get data %+v", resp)
	return &resp, nil
}

func (r *recordBiz) Layer(ctx context.Context, startTime time.Time, endTime time.Time) (*v1.ListTopologyLayerResponse, error) {
	hosts, err := r.ds.Records().List(ctx, "host", startTime, endTime)
	pods, err := r.ds.Records().List(ctx, "pod", startTime, endTime)
	containers, err := r.ds.Records().List(ctx, "container", startTime, endTime)

	if err != nil {
		if errors.Is(err, qmgo.ErrQueryResultTypeInconsistent) {
			return nil, errno.ErrRecordConnectFail
		}
	}
	var resp v1.ListTopologyLayerResponse
	log.Debug("[biz] get data from database successfully")
	var podToHost map[string]string = make(map[string]string)
	var topologyLayer model.APITopology = make(model.APITopology)
	for key, host := range hosts {
		topologyLayer[key] = model.APIHostTopology{NodeSummary: host, Children: make(map[string]model.APIPodTopology)}
	}
	for key, pod := range pods {
		for _, parent := range pod.Parents {
			if parent.TopologyID == "hosts" {
				host := parent.ID
				topologyLayer[host].Children[key] = model.APIPodTopology{NodeSummary: pod, Children: make(model.NodeSummaries)}
				podToHost[key] = host
			}
		}
	}

	for key, container := range containers {
		for _, parent := range container.Parents {
			if parent.TopologyID == "pods" {
				pod := parent.ID
				topologyLayer[podToHost[pod]].Children[pod].Children[key] = container
			}
		}
	}
	resp.Layer = topologyLayer
	if err != nil {
		log.Errorf("[biz] Fail to encode object %v", err)
		return nil, errno.ErrRecordConnectFail
	}
	log.Debugf("get data %+v", resp)
	return &resp, nil
}

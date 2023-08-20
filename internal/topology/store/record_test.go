package store

import (
	"ClusterWatcher/internal/pkg/model"
	"github.com/weaveworks/scope/report"
	"reflect"
	"testing"
	"time"
)

func TestMergeAll(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name  string
		input []model.RecordM
		want  model.RecordM
	}{
		{
			name: "merge add",
			input: []model.RecordM{
				{
					Timestamp: now,
					Nodes: model.NodeSummaries{
						"af1a69ee-9839-4bb7-881d-dea9a8a79d5a;\u003cpod\u003e": model.NodeSummary{
							ID:         "af1a69ee-9839-4bb7-881d-dea9a8a79d5a;\u003cpod\u003e",
							Label:      "kube-proxy-6drgt",
							LabelMinor: "1 container",
							Rank:       "kube-system/kube-proxy-6drgt",
							Shape:      "heptagon",
							Metadata: []report.MetadataRow{report.MetadataRow{
								ID:       "kubernetes_state",
								Label:    "State",
								Value:    "Running",
								Priority: 2.0,
							}, report.MetadataRow{
								ID:       "kubernetes_ip",
								Label:    "IP",
								Value:    "192.168.1.102",
								Priority: 3.0,
								Datatype: "ip",
							}, report.MetadataRow{
								ID:       "container",
								Label:    "# Containers",
								Value:    "1",
								Priority: 4.0,
								Datatype: "number",
							},
							},
							Parents: []model.Parent{model.Parent{
								ID:         "10b403b8-a710-4711-b353-af4cc3c30254;\u003cdaemonset\u003e",
								Label:      "kube-proxy",
								TopologyID: "kube-controllers",
							}, model.Parent{
								ID:         "worknode1;\u003chost\u003e",
								Label:      "worknode1",
								TopologyID: "hosts",
							},
							},
							Metrics: nil,
							Tables: []report.Table{
								report.Table{
									ID:      "kubernetes_labels_",
									Label:   "Kubernetes labels",
									Type:    "property-list",
									Columns: nil,
									Rows: []report.Row{
										report.Row{
											ID: "label_k8s-app",
											Entries: map[string]string{
												"label": "k8s-app", "value": "kube-proxy",
											},
										},
									},
								},
							},
							Adjacency: model.IDList{"30d1707a-13c8-48e8-87c6-d956b4d7ab48;\u003cpod\u003e"},
						},
					},
				},
				{
					Timestamp: now,
					Nodes: model.NodeSummaries{
						"af1a69ee-9839-4bb7-881d-dea9a8a79d5a;\u003cpod\u003e": model.NodeSummary{
							ID:         "af1a69ee-9839-4bb7-881d-dea9a8a79d5a;\u003cpod\u003e",
							Label:      "kube-proxy-6drgt",
							LabelMinor: "1 container",
							Rank:       "kube-system/kube-proxy-6drgt",
							Shape:      "heptagon",
							Metadata: []report.MetadataRow{report.MetadataRow{
								ID:       "kubernetes_state",
								Label:    "State",
								Value:    "Running",
								Priority: 2.0,
							}, {
								ID:       "kubernetes_ip",
								Label:    "IP",
								Value:    "192.168.1.103",
								Priority: 3.0,
								Datatype: "ip",
							}, {
								ID:       "container",
								Label:    "# Containers",
								Value:    "2",
								Priority: 4.0,
								Datatype: "number",
							},
							},
							Parents: []model.Parent{model.Parent{
								ID:         "10b403b8-a710-4711-b353-af4cc3c30254;\u003cdaemonset\u003e",
								Label:      "kube-proxy",
								TopologyID: "kube-controllers",
							}, {
								ID:         "worknode1;\u003chost\u003e",
								Label:      "worknode1",
								TopologyID: "hosts",
							},
							},
							Metrics: nil,
							Tables: []report.Table{
								{
									ID:      "kubernetes_labels_",
									Label:   "Kubernetes labels",
									Type:    "property-list",
									Columns: nil,
									Rows: []report.Row{
										report.Row{
											ID: "label_k8s-app",
											Entries: map[string]string{
												"label": "k8s-app", "value": "kube-proxy",
											},
										},
									},
								},
							},
							Adjacency: model.IDList{"9f6b0c19-9e75-4ccd-affb-6ccd937d3434;\u003cpod\u003e"},
						}, "cbb9698d-e577-4c1d-aa21-82100b11b591;\\u003cpersistent_volume\\u003e": model.NodeSummary{
							ID:         "af1a69ee-9839-4bb7-881d-dea9a8a79d5a;\u003cpod\u003e",
							Label:      "mongo-pv-storage",
							LabelMinor: "1 container",
							Rank:       "kube-system/kube-proxy-6drgt",
							Shape:      "heptagon",
							Metadata: []report.MetadataRow{report.MetadataRow{
								ID:       "kubernetes_state",
								Label:    "State",
								Value:    "Running",
								Priority: 2.0,
							}, {
								ID:       "kubernetes_ip",
								Label:    "IP",
								Value:    "192.168.1.102",
								Priority: 3.0,
								Datatype: "ip",
							}, {
								ID:       "container",
								Label:    "# Containers",
								Value:    "1",
								Priority: 4.0,
								Datatype: "number",
							},
							},
							Parents: []model.Parent{{
								ID:         "10b403b8-a710-4711-b353-af4cc3c30254;\u003cdaemonset\u003e",
								Label:      "kube-proxy",
								TopologyID: "kube-controllers",
							}, {
								ID:         "worknode1;\u003chost\u003e",
								Label:      "worknode1",
								TopologyID: "hosts",
							},
							},
							Metrics: nil,
							Tables: []report.Table{
								{
									ID:      "kubernetes_labels_",
									Label:   "Kubernetes labels",
									Type:    "property-list",
									Columns: nil,
									Rows: []report.Row{
										report.Row{
											ID: "label_k8s-app",
											Entries: map[string]string{
												"label": "k8s-app", "value": "kube-proxy",
											},
										},
									},
								},
							},
						},
					},
				},
			},
			want: model.RecordM{
				Timestamp: now,
				Nodes: model.NodeSummaries{
					"af1a69ee-9839-4bb7-881d-dea9a8a79d5a;\u003cpod\u003e": model.NodeSummary{
						ID:         "af1a69ee-9839-4bb7-881d-dea9a8a79d5a;\u003cpod\u003e",
						Label:      "kube-proxy-6drgt",
						LabelMinor: "1 container",
						Rank:       "kube-system/kube-proxy-6drgt",
						Shape:      "heptagon",
						Metadata: []report.MetadataRow{report.MetadataRow{
							ID:       "kubernetes_state",
							Label:    "State",
							Value:    "Running",
							Priority: 2.0,
						}, {
							ID:       "kubernetes_ip",
							Label:    "IP",
							Value:    "192.168.1.103",
							Priority: 3.0,
							Datatype: "ip",
						}, {
							ID:       "container",
							Label:    "# Containers",
							Value:    "2",
							Priority: 4.0,
							Datatype: "number",
						},
						},
						Parents: []model.Parent{model.Parent{
							ID:         "10b403b8-a710-4711-b353-af4cc3c30254;\u003cdaemonset\u003e",
							Label:      "kube-proxy",
							TopologyID: "kube-controllers",
						}, {
							ID:         "worknode1;\u003chost\u003e",
							Label:      "worknode1",
							TopologyID: "hosts",
						},
						},
						Metrics: nil,
						Tables: []report.Table{
							report.Table{
								ID:      "kubernetes_labels_",
								Label:   "Kubernetes labels",
								Type:    "property-list",
								Columns: nil,
								Rows: []report.Row{
									report.Row{
										ID: "label_k8s-app",
										Entries: map[string]string{
											"label": "k8s-app", "value": "kube-proxy",
										},
									},
								},
							},
						},
						Adjacency: model.IDList{"30d1707a-13c8-48e8-87c6-d956b4d7ab48;\u003cpod\u003e", "9f6b0c19-9e75-4ccd-affb-6ccd937d3434;\u003cpod\u003e"},
					}, "cbb9698d-e577-4c1d-aa21-82100b11b591;\\u003cpersistent_volume\\u003e": model.NodeSummary{
						ID:         "af1a69ee-9839-4bb7-881d-dea9a8a79d5a;\u003cpod\u003e",
						Label:      "mongo-pv-storage",
						LabelMinor: "1 container",
						Rank:       "kube-system/kube-proxy-6drgt",
						Shape:      "heptagon",
						Metadata: []report.MetadataRow{report.MetadataRow{
							ID:       "kubernetes_state",
							Label:    "State",
							Value:    "Running",
							Priority: 2.0,
						}, {
							ID:       "kubernetes_ip",
							Label:    "IP",
							Value:    "192.168.1.102",
							Priority: 3.0,
							Datatype: "ip",
						}, {
							ID:       "container",
							Label:    "# Containers",
							Value:    "1",
							Priority: 4.0,
							Datatype: "number",
						},
						},
						Parents: []model.Parent{model.Parent{
							ID:         "10b403b8-a710-4711-b353-af4cc3c30254;\u003cdaemonset\u003e",
							Label:      "kube-proxy",
							TopologyID: "kube-controllers",
						}, {
							ID:         "worknode1;\u003chost\u003e",
							Label:      "worknode1",
							TopologyID: "hosts",
						},
						},
						Metrics: nil,
						Tables: []report.Table{
							{
								ID:      "kubernetes_labels_",
								Label:   "Kubernetes labels",
								Type:    "property-list",
								Columns: nil,
								Rows: []report.Row{
									report.Row{
										ID: "label_k8s-app",
										Entries: map[string]string{
											"label": "k8s-app", "value": "kube-proxy",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeAll(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeAll() = \n%v, \nwant %v", got, tt.want)
			}
		})
	}
}

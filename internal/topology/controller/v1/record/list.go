package record

import (
	"ClusterWatcher/internal/pkg/core"
	"ClusterWatcher/internal/pkg/log"
	v1 "ClusterWatcher/pkg/api/topology/v1"
	"github.com/gin-gonic/gin"
)

func (ctrl *RecordController) List(c *gin.Context) {
	log.Infow("List record function called.")

	var r v1.ListRecordRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}
	name := c.Param("name")
	resp, err := ctrl.b.Records().List(c, name, r.StartTime, r.EndTime)

	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

func (ctrl *RecordController) GetName(c *gin.Context) {
	log.Infow("List record function called.")

	var r v1.ListTopologyNameRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	resp, err := ctrl.b.Records().GetName(c, r.StartTime, r.EndTime)

	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

func (ctrl *RecordController) GetNs(c *gin.Context) {
	log.Infow("List record function called.")

	var r v1.ListTopologyNsRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	resp, err := ctrl.b.Records().GetNs(c, r.StartTime, r.EndTime)

	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

func (ctrl *RecordController) Layer(c *gin.Context) {
	log.Infow("List record function called.")

	var r v1.ListTopologyLayerRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	resp, err := ctrl.b.Records().Layer(c, r.StartTime, r.EndTime)

	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

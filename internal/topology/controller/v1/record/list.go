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

func (ctrl *RecordController) Group(c *gin.Context) {
	log.Infow("List record function called.")

	var r v1.ListTopologyGroupRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	resp, err := ctrl.b.Records().Group(c, r.StartTime, r.EndTime)

	if err != nil {
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
}

func (ctrl *RecordController) Layer(c *gin.Context) {
	log.Infow("List record function called.")

	var r v1.ListTopologyGroupRequest
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

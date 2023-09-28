package v1

import (
	"ClusterWatcher/internal/pkg/model"
	"time"
)

type ListRecordResponse struct {
	Records model.NodeSummaries `json:"records"`
}

type ListRecordRequest struct {
	StartTime time.Time `form:"start_time" binding:"required" time_format:"2006-01-02T15:04:05Z"`
	EndTime   time.Time `form:"end_time" binding:"required" time_format:"2006-01-02T15:04:05Z"`
}

type ListTopologyNameResponse struct {
	Topology model.APITopologyGroup `json:"topologyGroup"`
}

type ListTopologyNameRequest struct {
	StartTime time.Time `form:"start_time" binding:"required" time_format:"2006-01-02T15:04:05Z"`
	EndTime   time.Time `form:"end_time" binding:"required" time_format:"2006-01-02T15:04:05Z"`
}

type ListTopologyNsResponse struct {
	Topology model.APITopologyGroup `json:"topologyGroup"`
}

type ListTopologyNsRequest struct {
	StartTime time.Time `form:"start_time" binding:"required" time_format:"2006-01-02T15:04:05Z"`
	EndTime   time.Time `form:"end_time" binding:"required" time_format:"2006-01-02T15:04:05Z"`
}

type ListTopologyLayerResponse struct {
	Layer model.APITopology `json:"topologyLayer"`
}

type ListTopologyLayerRequest struct {
	StartTime time.Time `form:"start_time" binding:"required" time_format:"2006-01-02T15:04:05Z"`
	EndTime   time.Time `form:"end_time" binding:"required" time_format:"2006-01-02T15:04:05Z"`
}

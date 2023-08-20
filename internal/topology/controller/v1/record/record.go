package record

import (
	"ClusterWatcher/internal/topology/biz"
	"ClusterWatcher/internal/topology/store"
)

type RecordController struct {
	b biz.IBiz
}

func New(ds store.IStore) *RecordController {
	return &RecordController{b: biz.NewBiz(ds)}
}

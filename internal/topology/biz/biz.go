package biz

import (
	"ClusterWatcher/internal/topology/biz/record"

	"ClusterWatcher/internal/topology/store"
)

type IBiz interface {
	Records() record.RecordBiz
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// biz 是 IBiz 的一个具体实现.
type biz struct {
	ds store.IStore
}

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}

// Records 返回一个实现了 UserBiz 接口的实例.
func (b *biz) Records() record.RecordBiz {
	return record.New(b.ds)
}

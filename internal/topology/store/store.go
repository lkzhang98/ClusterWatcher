package store

import (
	"github.com/qiniu/qmgo"
	"sync"
)

var (
	once sync.Once
	// 全局变量，方便其它包直接调用已初始化好的 S 实例.
	S *datastore
)

// IStore 定义了 Store 层需要实现的方法.
type IStore interface {
	MongoDB() *qmgo.Database
	Records() RecordStore
}

// datastore 是 IStore 的一个具体实现.
type datastore struct {
	mongo *qmgo.Database
}

// 确保 datastore 实现了 IStore 接口.
var _ IStore = (*datastore)(nil)

// NewStore 创建一个 IStore 类型的实例.
func NewStore(mongo *qmgo.Database) *datastore {
	// 确保 S 只被初始化一次
	once.Do(func() {
		S = &datastore{mongo}
	})

	return S
}

func (ds *datastore) MongoDB() *qmgo.Database {
	return ds.mongo
}

func (ds *datastore) Records() RecordStore {
	return newRecords(ds.mongo)
}

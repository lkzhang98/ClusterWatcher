package collecter

import (
	"ClusterWatcher/internal/pkg/model"
	codec "github.com/ugorji/go/codec"
	"time"
)

type Item interface {
	serializer() *model.RecordM
}

type item struct {
	Timestamp time.Time `json:"timestamp"`
	Data      []byte    `json:"data"`
}

var _ Item = &item{}

func NewItem(timestamp time.Time, data []byte) *item {
	return &item{
		Timestamp: timestamp,
		Data:      data,
	}
}

// 当解析出错时，能够从其中恢复这个问题，或者是不是字节流丢失的比较多
func (i *item) serializer() *model.RecordM {
	var records model.Record

	_ = codec.NewDecoderBytes(i.Data, &codec.JsonHandle{}).Decode(&records)
	record := &model.RecordM{
		Timestamp: i.Timestamp,
		Nodes:     records.Nodes,
	}

	return record
}

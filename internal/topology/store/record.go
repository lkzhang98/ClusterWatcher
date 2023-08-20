package store

import (
	"ClusterWatcher/internal/pkg/log"
	"ClusterWatcher/internal/pkg/model"
	"context"
	"github.com/qiniu/qmgo"
	"github.com/weaveworks/scope/render/detailed"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type RecordStore interface {
	List(context.Context, string, time.Time, time.Time) (model.NodeSummaries, error)
}

// RecordStore 接口的实现.
type records struct {
	database *qmgo.Database
}

type Record detailed.NodeSummaries

// 确保 records 实现了 RecordStore 接口.
var _ RecordStore = (*records)(nil)

func newRecords(database *qmgo.Database) *records {
	return &records{database}
}

func (r *records) List(ctx context.Context, name string, startTime time.Time, endTime time.Time) (model.NodeSummaries, error) {

	log.Debug("store search time:", startTime.UTC(), endTime.Format(time.RFC3339))

	filter := bson.M{
		"timestamp": bson.M{
			"$gte": startTime.UTC(),
			"$lte": endTime.UTC(),
		},
	}
	var recordList RecordList
	err := r.database.Collection(name).Find(context.TODO(), filter).All(&recordList.Records)

	if err != nil {
		log.Errorf("[Store] get records from database error: %v", err)
		return model.NodeSummaries{}, err
	}

	log.Debugf("store count: %d", len(recordList.Records))

	return recordList.render(), nil
}

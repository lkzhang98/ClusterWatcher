package store

import (
	"ClusterWatcher/internal/pkg/log"
	"ClusterWatcher/internal/pkg/model"
	"context"
	"github.com/go-redis/redis"
	"github.com/qiniu/qmgo"
	"github.com/ugorji/go/codec"
	"github.com/weaveworks/scope/render/detailed"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type RecordStore interface {
	List(context.Context, string, time.Time, time.Time) (model.NodeSummaries, error)
}

// RecordStore 接口的实现.
type records struct {
	database    *qmgo.Database
	redisClient *redis.Client
}

type Record detailed.NodeSummaries

// 确保 records 实现了 RecordStore 接口.
var _ RecordStore = (*records)(nil)

func newRecords(database *qmgo.Database, client *redis.Client) *records {
	return &records{database, client}
}

func (r *records) List(ctx context.Context, name string, startTime time.Time, endTime time.Time) (model.NodeSummaries, error) {

	log.Debug("store search time:", startTime.UTC(), endTime.Format(time.RFC3339))

	key := startTime.String() + endTime.String() + name

	data, err := r.redisClient.Get(key).Result()
	if err != nil {
		log.Errorf("[Store] get records from redis error: %v", err)
	}
	if data != "" {
		log.Info("[store] get records from redis")
		var record model.NodeSummaries
		dec := codec.NewDecoderBytes([]byte(data), &codec.BincHandle{})
		err = dec.Decode(&record)

		if err != nil {
			log.Errorf("[store] decode error: %v", err)
		} else {
			return record, nil
		}
	}

	filter := bson.M{
		"timestamp": bson.M{
			"$gte": startTime.UTC(),
			"$lte": endTime.UTC(),
		},
	}
	var recordList RecordList
	err = r.database.Collection(name).Find(context.TODO(), filter).All(&recordList.Records)

	if err != nil {
		log.Errorf("[Store] get records from database error: %v", err)
		return model.NodeSummaries{}, err
	}

	log.Debugf("store count: %d", len(recordList.Records))
	result := recordList.render()
	var bytes []byte
	encoder := codec.NewEncoderBytes(&bytes, &codec.BincHandle{})
	err = encoder.Encode(result)
	log.Errorf("[store] encode error: %v", err)
	err = r.redisClient.Set(key, string(bytes), 0).Err()
	if err != nil {
		log.Errorf("[store] redis set data error: %v", err)
	}
	log.Info("[store] redis set data successfully!")
	return result, nil
}

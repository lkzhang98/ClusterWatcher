package model

import (
	"time"
)

type Record struct {
	Nodes NodeSummaries `json:"nodes"`
}

type RecordM struct {
	Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
	Nodes     NodeSummaries `json:"nodes" bson:"nodes"`
}

func (r *RecordM) GetCollectionName() string {
	return "Record"
}

package db

import (
	"context"
	"github.com/qiniu/qmgo"

	"ClusterWatcher/internal/pkg/log"
	"strconv"
)

type MongodbOptions struct {
	Host       string
	Port       int
	Database   string
	Collection string
}

func (o *MongodbOptions) DSN() string {
	p := strconv.Itoa(o.Port)
	return "mongodb://" + o.Host + ":" + p
}

func NewMongoDB(opts *MongodbOptions) (*qmgo.Database, error) {
	cli, err := qmgo.NewClient(context.TODO(), &qmgo.Config{Uri: opts.DSN()})

	if err != nil {
		log.Fatal("Mongo init failed!")
		return nil, err
	}
	database := cli.Database(opts.Database)
	return database, nil
}

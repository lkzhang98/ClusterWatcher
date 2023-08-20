package consumer

import (
	"ClusterWatcher/internal/pkg/log"
	"ClusterWatcher/internal/pkg/model"
	"context"
	"fmt"
	kelleyRabbimqPool "gitee.com/tym_hmm/rabbitmq-pool-go"
	"github.com/qiniu/qmgo"
	"github.com/ugorji/go/codec"
)

var bh codec.BincHandle

func ConsumeMessage(instanceConsumePool *kelleyRabbimqPool.RabbitPool, database *qmgo.Database) {
	log.Debug("ConsumeMessage called")
	consumeHostMessage := &kelleyRabbimqPool.ConsumeReceive{
		ExchangeName: "topologydata", //队列名称
		ExchangeType: kelleyRabbimqPool.EXCHANGE_TYPE_DIRECT,
		Route:        "/" + "host",
		QueueName:    "host",
		IsTry:        true,  //是否重试
		IsAutoAck:    false, //是否自动确认消息
		MaxReTry:     5,     //最大重试次数
		EventFail: func(code int, e error, data []byte) {
			log.Error("Handle host message error: %v", e)
		},
		EventSuccess: func(data []byte, header map[string]interface{}, retryClient kelleyRabbimqPool.RetryClientInterface) bool { //如果返回true 则无需重试
			_ = retryClient.Ack() //确认消息
			fmt.Println("host message handle")
			var record model.RecordM
			dec := codec.NewDecoderBytes(data, &bh)
			err := dec.Decode(&record)
			if err != nil {
				log.Error("host message decode error : %v", err)
				return false
			}
			collection := database.Collection("host")
			_, err = collection.InsertOne(context.TODO(), record)
			if err != nil {
				log.Error("host message storage error : %v", err)
				return false
			}
			return true
		},
	}

	consumePodMessage := &kelleyRabbimqPool.ConsumeReceive{
		ExchangeName: "topologydata", //队列名称
		ExchangeType: kelleyRabbimqPool.EXCHANGE_TYPE_DIRECT,
		Route:        "/" + "pod",
		QueueName:    "pod",
		IsTry:        true,  //是否重试
		IsAutoAck:    false, //是否自动确认消息
		MaxReTry:     5,     //最大重试次数
		EventFail: func(code int, e error, data []byte) {
			log.Error("Handle host message error: %v", e)
		},
		EventSuccess: func(data []byte, header map[string]interface{}, retryClient kelleyRabbimqPool.RetryClientInterface) bool { //如果返回true 则无需重试
			_ = retryClient.Ack() //确认消息
			fmt.Println("host message handle")
			var record model.RecordM
			dec := codec.NewDecoderBytes(data, &bh)
			err := dec.Decode(&record)
			if err != nil {
				log.Error("host message decode error : %v", err)
				return false
			}
			collection := database.Collection("pod")
			_, err = collection.InsertOne(context.TODO(), record)
			if err != nil {
				log.Error("host message storage error : %v", err)
				return false
			}
			return true
		},
	}

	consumeContainerMessage := &kelleyRabbimqPool.ConsumeReceive{
		ExchangeName: "topologydata", //队列名称
		ExchangeType: kelleyRabbimqPool.EXCHANGE_TYPE_DIRECT,
		Route:        "/" + "container",
		QueueName:    "container",
		IsTry:        true,  //是否重试
		IsAutoAck:    false, //是否自动确认消息
		MaxReTry:     5,     //最大重试次数
		EventFail: func(code int, e error, data []byte) {
			log.Error("Handle host message error: %v", e)
		},
		EventSuccess: func(data []byte, header map[string]interface{}, retryClient kelleyRabbimqPool.RetryClientInterface) bool { //如果返回true 则无需重试
			_ = retryClient.Ack() //确认消息
			fmt.Println("host message handle")
			var record model.RecordM
			dec := codec.NewDecoderBytes(data, &bh)
			err := dec.Decode(&record)
			if err != nil {
				log.Error("host message decode error : %v", err)
				return false
			}
			collection := database.Collection("container")
			_, err = collection.InsertOne(context.TODO(), record)
			if err != nil {
				log.Error("host message storage error : %v", err)
				return false
			}
			return true
		},
	}

	instanceConsumePool.RegisterConsumeReceive(consumeHostMessage)
	instanceConsumePool.RegisterConsumeReceive(consumePodMessage)
	instanceConsumePool.RegisterConsumeReceive(consumeContainerMessage)

	err := instanceConsumePool.RunConsume()
	if err != nil {
		fmt.Println(err)
	}
}

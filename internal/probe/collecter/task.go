package collecter

import (
	"ClusterWatcher/internal/pkg/log"
	pool "gitee.com/tym_hmm/rabbitmq-pool-go"
	"github.com/ugorji/go/codec"
	"io/ioutil"
	"net/http"
	"time"
)

type CollectFunc func(url string, channel *pool.RabbitPool, chName string)

type CollectConfig struct {
	Url    string
	ChName string
}

type CollectConfigs map[string]CollectConfig

var bh codec.BincHandle

func CollectTask() CollectFunc {
	return func(url string, channel *pool.RabbitPool, chName string) {
		location, _ := time.LoadLocation("Asia/Shanghai")
		current := time.Now().In(location)
		response, err := http.Get(url)
		if err != nil {
			log.Errorf("grab api error: %v", err)
		}

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Errorf("read data from response error : %v", err)
		}
		record := NewItem(current, []byte(body)).serializer()
		var bytes []byte
		encoder := codec.NewEncoderBytes(&bytes, &bh)
		err = encoder.Encode(record)
		if err != nil {
			log.Errorf("encode object error : %v", err)
		}

		report := pool.GetRabbitMqDataFormat("topologydata", pool.EXCHANGE_TYPE_DIRECT, chName, "/"+chName, string(bytes))
		log.Debugf("send message to %s", chName)
		mqError := channel.Push(report)
		if err != nil {
			log.Errorf("Failed to publish message: %v", mqError.Message)
		}
	}
}

package collecter

//
//import (
//	"ClusterWatcher/internal/pkg/log"
//	"context"
//	"encoding/json"
//	pool "gitee.com/tym_hmm/rabbitmq-pool-go"
//	"github.com/hibiken/asynq"
//	"io/ioutil"
//	"net/http"
//	"time"
//)
//
//const (
//	CollectHost      = "collect:host"
//	CollectPod       = "collect:pod"
//	CollectContainer = "collect:container"
//)
//
//func NewCollectHostTask(url string, channel *pool.RabbitPool) *asynq.Task {
//	payload := map[string]interface{}{"url": url, "channel": channel}
//	bytes, err := json.Marshal(payload)
//	if err != nil {
//		log.Errorf("Task payload error: %s", "CollectHostTask")
//	}
//	return asynq.NewTask(CollectHost, bytes)
//}
//
//func NewCollectPodTask(url string, channel *pool.RabbitPool) *asynq.Task {
//	payload := map[string]interface{}{"url": url, "channel": channel}
//	bytes, err := json.Marshal(payload)
//	if err != nil {
//		log.Errorf("Task payload error: %s", "CollectPodTask")
//	}
//	return asynq.NewTask(CollectPod, bytes)
//}
//
//func NewCollectContainerTask(url string, channel *pool.RabbitPool) *asynq.Task {
//	payload := map[string]interface{}{"url": url, "channel": channel}
//	bytes, err := json.Marshal(payload)
//	if err != nil {
//		log.Errorf("Task payload error: %s", "CollectContainerTask")
//	}
//	return asynq.NewTask(CollectContainer, bytes)
//}
//
//func HandleCollectHostTask(ctx context.Context, t *asynq.Task) error {
//	var data map[string]interface{}
//	err := json.Unmarshal(t.Payload(), &data)
//	if err != nil {
//		log.Errorf("Task payload error: %s", err)
//		return err
//	}
//	location, _ := time.LoadLocation("Asia/Shanghai")
//	current := time.Now().In(location)
//	response, err := http.Get(data["url"].(string))
//	if err != nil {
//		log.Errorf("grab api error: %v", err)
//		return err
//	}
//
//	defer response.Body.Close()
//
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Errorf("read data from response error : %v", err)
//		return err
//	}
//	bytes := NewItem(current, []byte(body)).serializer()
//	report := pool.GetRabbitMqDataFormat("topologydata", pool.EXCHANGE_TYPE_DIRECT, "host", "/", string(bytes))
//	err = data["channel"].(*pool.RabbitPool).Push(report)
//	if err != nil {
//		log.Fatalf("Failed to publish message: %v", err)
//	}
//	return nil
//}
//
//func HandleCollectPodTask(ctx context.Context, t *asynq.Task) error {
//	var data map[string]interface{}
//	err := json.Unmarshal(t.Payload(), &data)
//	if err != nil {
//		log.Errorf("Task payload error: %s", err)
//		return err
//	}
//	location, _ := time.LoadLocation("Asia/Shanghai")
//	current := time.Now().In(location)
//	response, err := http.Get(data["url"].(string))
//	if err != nil {
//		log.Errorf("grab api error: %v", err)
//		return err
//	}
//
//	defer response.Body.Close()
//
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Errorf("read data from response error : %v", err)
//		return err
//	}
//	bytes := NewItem(current, []byte(body)).serializer()
//	report := pool.GetRabbitMqDataFormat("topologydata", pool.EXCHANGE_TYPE_DIRECT, "pod", "/", string(bytes))
//	err = data["channel"].(*pool.RabbitPool).Push(report)
//	if err != nil {
//		log.Fatalf("Failed to publish message: %v", err)
//	}
//	return nil
//}
//
//func HandleCollectContainerTask(ctx context.Context, t *asynq.Task) error {
//	var data map[string]interface{}
//	err := json.Unmarshal(t.Payload(), &data)
//	if err != nil {
//		log.Errorf("Task payload error: %s", err)
//		return err
//	}
//	location, _ := time.LoadLocation("Asia/Shanghai")
//	current := time.Now().In(location)
//	response, err := http.Get(data["url"].(string))
//	if err != nil {
//		log.Errorf("grab api error: %v", err)
//		return err
//	}
//
//	defer response.Body.Close()
//
//	body, err := ioutil.ReadAll(response.Body)
//	if err != nil {
//		log.Errorf("read data from response error : %v", err)
//		return err
//	}
//	bytes := NewItem(current, []byte(body)).serializer()
//	report := pool.GetRabbitMqDataFormat("topologydata", pool.EXCHANGE_TYPE_DIRECT, "container", "/", string(bytes))
//	err = data["channel"].(*pool.RabbitPool).Push(report)
//	if err != nil {
//		log.Fatalf("Failed to publish message: %v", err)
//	}
//	return nil
//}

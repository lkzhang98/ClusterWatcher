package probe

import (
	"ClusterWatcher/internal/pkg/db"
	"ClusterWatcher/internal/pkg/log"
	"ClusterWatcher/internal/probe/collecter"
	kelleyRabbimqPool "gitee.com/tym_hmm/rabbitmq-pool-go"
	"github.com/qiniu/qmgo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var oncePool sync.Once
var instanceRPool *kelleyRabbimqPool.RabbitPool
var instanceCPool *kelleyRabbimqPool.RabbitPool

const (
	recommendedHomeDir = ".probe"
	defaultConfigName  = "probe"
)

var prefix string
var cfgFile string

func initConfig() {
	if cfgFile != "" {
		// 从命令行选项指定的配置文件中读取
		viper.SetConfigFile(cfgFile)
	} else {
		// 查找用户主目录
		home, err := os.UserHomeDir()
		// 如果获取用户主目录失败，打印 `'Error: xxx` 错误，并退出程序（退出码为 1）
		cobra.CheckErr(err)

		// 将用 `$HOME/<recommendedHomeDir>` 目录加入到配置文件的搜索路径中
		viper.AddConfigPath(filepath.Join(home, recommendedHomeDir))

		// 把当前目录加入到配置文件的搜索路径中
		viper.AddConfigPath("../../configs/")

		// 设置配置文件格式为 YAML (YAML 格式清晰易读，并且支持复杂的配置结构)
		viper.SetConfigType("yaml")

		// 配置文件名称（没有文件扩展名）
		viper.SetConfigName(defaultConfigName)
	}

	// 读取匹配的环境变量
	viper.AutomaticEnv()

	// 读取环境变量的前缀为 MINIBLOG，如果是 miniblog，将自动转变为大写。
	viper.SetEnvPrefix("PROBE")

	// 以下 2 行，将 viper.Get(key) key 字符串中 '.' 和 '-' 替换为 '_'
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		log.Errorw("Failed to read viper configuration file", "err", err)
	}
	prefix = viper.GetString("runmode") + "."
	// 打印 viper 当前使用的配置文件，方便 Debug.
	log.Debugw("Using configs file", "file", viper.ConfigFileUsed())
}

func logOptions() *log.Options {
	return &log.Options{
		DisableCaller:     viper.GetBool(prefix + "log.disable-caller"),
		DisableStacktrace: viper.GetBool(prefix + "log.disable-stacktrace"),
		Level:             viper.GetString(prefix + "log.level"),
		Format:            viper.GetString(prefix + "log.format"),
		OutputPaths:       viper.GetStringSlice(prefix + "log.output-paths"),
	}
}

func initProduceRabbitmq() *kelleyRabbimqPool.RabbitPool {
	oncePool.Do(func() {
		//初始化生产者
		instanceRPool = kelleyRabbimqPool.NewProductPool()
		//初始化消费者
		//使用默认虚拟host "/"
		err := instanceRPool.Connect(viper.GetString(prefix+"rabbitmq.host"), viper.GetInt(prefix+"rabbitmq.port"), viper.GetString(prefix+"rabbitmq.user"), viper.GetString(prefix+"rabbitmq.password"))

		if err != nil {
			log.Error("rabbitmq pool create failed!")
		}
	})
	log.Info("rabbitmq pool create successfully!")
	return instanceRPool
}

func initConsumeRabbitmq() *kelleyRabbimqPool.RabbitPool {
	oncePool.Do(func() {
		//初始化生产者
		instanceCPool = kelleyRabbimqPool.NewConsumePool()
		//初始化消费者
		//使用默认虚拟host "/"
		err := instanceCPool.Connect(viper.GetString(prefix+"rabbitmq.host"), viper.GetInt(prefix+"rabbitmq.port"), viper.GetString(prefix+"rabbitmq.user"), viper.GetString(prefix+"rabbitmq.password"))
		if err != nil {
			log.Error("rabbitmq pool create failed!")
		}
	})
	log.Info("rabbitmq pool create successfully!")
	return instanceCPool
}

//func initRabbitMqConn() {
//	conn, err = amqp.Dial("amqp://" + viper.GetString("rabbitmq.user"):admin@ip:5672/")
//}

func initStore() *qmgo.Database {
	dbOptions := &db.MongodbOptions{
		Host:     viper.GetString(prefix + "mongodb.host"),
		Port:     viper.GetInt(prefix + "mongodb.port"),
		Database: viper.GetString(prefix + "mongodb.database"),
	}

	database, err := db.NewMongoDB(dbOptions)
	if err != nil {
		log.Fatalf("mongodb init failed!")
		os.Exit(1)
	}
	return database
}

func initTaskConfigs() collecter.CollectConfigs {
	taskConfigs := map[string]collecter.CollectConfig{
		viper.GetString("task.host.name"):      {viper.GetString("task.host.url"), viper.GetString("task.host.channel-name")},
		viper.GetString("task.pod.name"):       {viper.GetString("task.pod.url"), viper.GetString("task.pod.channel-name")},
		viper.GetString("task.container.name"): {viper.GetString("task.container.url"), viper.GetString("task.container.channel-name")},
	}
	return taskConfigs
}

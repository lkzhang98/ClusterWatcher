package topology

import (
	"ClusterWatcher/internal/pkg/db"
	"ClusterWatcher/internal/pkg/log"
	"ClusterWatcher/internal/topology/store"
	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

const (
	recommendedHomeDir = ".topology"
	defaultConfigName  = "topology"
)

var prefix string

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

func initStore() error {
	dbOptions := &db.MongodbOptions{
		Host:     viper.GetString(prefix + "mongodb.host"),
		Port:     viper.GetInt(prefix + "mongodb.port"),
		Database: viper.GetString(prefix + "mongodb.database"),
	}

	cli, err := db.NewMongoDB(dbOptions)
	if err != nil {
		log.Fatalf("mongodb init failed!")
		return err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString(prefix+"redis.host") + ":" + viper.GetString(prefix+"redis.port"),
		Password: "",
		DB:       viper.GetInt(prefix + "redis.db"),
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		log.Fatalf("redis init failed! %v", err)
	}

	_ = store.NewStore(cli, rdb)
	return nil
}

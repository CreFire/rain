package config

import (
	"fmt"
	"github.com/CreFire/rain/model"
	"github.com/CreFire/rain/pkg/log"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func IsDev() bool {
	return GlobalConfig.Model == "development"
}
var GlobalConfig *model.Config

// config init
func init() {
	logger := log.Default()
	GlobalConfig = &model.Config{}
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./conf/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Error("viper ReadConfig err", log.Any("err", any(err)))
		return
	}
	err = viper.Unmarshal(&GlobalConfig)
	if err != nil {
		logger.Error("viper ReadConfig err", log.Any("err", any(err)),log.Any("config",GlobalConfig))
		return
	}
	logger.Info("config init",log.Any("config",GlobalConfig))
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
	viper.WatchConfig()
	viper.SetDefault("rain.LogDir", "./log")
}

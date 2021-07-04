package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	var err error
	config = viper.New()
	// 配置文件类型
	config.SetConfigType("yaml")
	// 文件名
	config.SetConfigName(env)
	// 文件路径
	config.AddConfigPath("./config/")
	// 读取配置文件
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
}

func GetConfig() *viper.Viper {
	return config
}

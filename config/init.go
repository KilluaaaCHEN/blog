package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

/**
初始化配置文件
*/
func InitConfig() {

	viper.SetConfigFile("config/app.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	//配置文件热更新
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed:%v \n", e.Name)
	})

}

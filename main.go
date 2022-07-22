package main

import (
	. "awesomeSystem/routers"
	"awesomeSystem/utils/logger"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var logConf = new(logger.LogConfig)

func main() {
	// viper
	viper.SetConfigFile("./config/logconf.json") //指定json配置文件的路径
	err := viper.ReadInConfig()                  // 读取配置信息
	if err != nil {                              // 配置失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err = viper.Unmarshal(logConf); err != nil {
		panic(fmt.Errorf("unmarshal conf failed, err:%s \n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()

	// zap 初始化
	if err := logger.InitLogger(logConf); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	if err := R.Run(":8877"); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

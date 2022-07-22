package main

import (
	. "awesomeSystem/app/weight/router"
	"awesomeSystem/global"
	"awesomeSystem/initialize"
	"fmt"
	"github.com/fatih/color"
	"log"
)

//var logConf = new(logger.LogConfig)

func main() {
	// 1.初始化yaml配置
	initialize.InitConfig()
	// 2.初始化日志信息
	initialize.InitLogger()
	// 3.初始化mongodb
	initialize.InitMongodb()
	// 4.初始化mysql
	initialize.InitMySQL()
	// 5.初始化casbin
	initialize.InitACS()
	// 6.初始化redis
	initialize.InitRedis()

	color.Cyan("go-gin服务开始了")

	if err := R.Run(fmt.Sprintf(":%d", global.Settings.WeightInfo.Port)); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}

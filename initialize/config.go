package initialize

import (
	"awesomeSystem/config"
	"awesomeSystem/global"
	"github.com/fatih/color"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitConfig() {
	// 实例化viper
	v := viper.New()
	//文件的路径如何设置
	v.SetConfigFile("./settings.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	//给serverConfig初始值
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	zap.L().Info("config loaded!", zap.String("config", "--------->loaded"))
	// 传递给全局变量
	global.Settings = serverConfig

	color.Blue("log file", global.Settings.LogsAddress)

}

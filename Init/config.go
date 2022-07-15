package Init

import (
	"dousheng/config"
	"dousheng/global"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

//完成viper的初始化
func Config() {
	// 实例化viper
	v := viper.New()
	//读取settings文件
	v.SetConfigFile("./settings.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//给serverConfig初始值
	serverConfig := config.ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	// 传递给全局变量
	global.Settings = serverConfig
	color.Green("Viper Config Initialize")
}

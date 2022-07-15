package main

import (
	"dousheng/Init"
	"dousheng/global"
	"fmt"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Router *gin.Engine

func main() {
	color.Cyan("dousheng server start!!!!")
	//启动gin,并配置端口,global.Settings.Port是yaml配置过的
	err := Router.Run(fmt.Sprintf(":%d", global.Settings.Port))
	if err != nil {
		zap.L().Info("this is server func", zap.String("error", "启动错误!"))
	}
}

func init() {
	Init.Config()
	Router = Init.Routers()
}

package Init

import (
	"dousheng/middlewares"
	"dousheng/router"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

//定义中间件及路由分组的配置
func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	ApiGroup := r.Group("/douyin/")
	router.User(ApiGroup)
	router.Feed(ApiGroup)
	color.Green("Routers Initialize")
	return r
}

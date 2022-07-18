package router

import (
	"dousheng/controller"

	"github.com/gin-gonic/gin"
)

func Feed(Router *gin.RouterGroup) {
	r := Router.Group("feed")
	{
		r.GET("", func(c *gin.Context) { controller.Feed(c) })
	}
}

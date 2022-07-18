package router

import (
	"dousheng/controller"

	"github.com/gin-gonic/gin"
)

func Publish(Router *gin.RouterGroup) {
	r := Router.Group("publish")
	{
		r.POST("action/", func(c *gin.Context) { controller.PublishAction(c) })

		r.GET("list/", func(c *gin.Context) { controller.PublishList(c) })
	}
}

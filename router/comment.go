package router

import (
	"dousheng/controller"

	"github.com/gin-gonic/gin"
)

func Comment(Router *gin.RouterGroup) {
	r := Router.Group("comment")
	{
		r.POST("action/", func(c *gin.Context) { controller.CommentAction(c) })

		r.GET("list/", func(c *gin.Context) { controller.CommentList(c) })
	}
}

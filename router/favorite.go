package router

import (
	"dousheng/controller"

	"github.com/gin-gonic/gin"
)

func Favorite(Router *gin.RouterGroup) {
	r := Router.Group("favorite")
	{
		r.POST("action/", func(c *gin.Context) { controller.FavoriteAction(c) })

		r.GET("list/", func(c *gin.Context) { controller.FavoriteList(c) })
	}

}

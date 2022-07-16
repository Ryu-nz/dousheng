package router

import (
	"dousheng/service"

	"github.com/gin-gonic/gin"
)

func Feed(Router *gin.RouterGroup) {
	r := Router.Group("feed")
	{
		r.GET("", func(c *gin.Context) {
			service.GetVedioList(c)
		})
	}
}

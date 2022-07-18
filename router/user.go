package router

import (
	"dousheng/controller"

	"github.com/gin-gonic/gin"
)

//用户接口路由
func User(Router *gin.RouterGroup) {
	r := Router.Group("user")
	{
		r.GET("/", func(c *gin.Context) { controller.GetUser(c) })

		r.POST("register/", func(c *gin.Context) { controller.UserRegister(c) })

		r.POST("login/", func(c *gin.Context) { controller.UserLogin(c) })
	}
}

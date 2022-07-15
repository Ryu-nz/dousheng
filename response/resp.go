package response

import "github.com/gin-gonic/gin"

// 全局返回
func Msg(c *gin.Context, httpCode int, code int, msg string, jsonStr interface{}) {
	c.JSON(httpCode, map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": jsonStr,
	})
	return
}

//错误返回
func Err(c *gin.Context) {}

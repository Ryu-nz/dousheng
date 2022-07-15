package utils

import (
	"dousheng/middlewares"
	"dousheng/response"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(c *gin.Context, UserID int, UserName string, Role int) string {
	//生成token信息
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		UserID:   UserID,
		UserName: UserName,
		Role:     Role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			// TODO 设置token过期时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //token -->30天过期
			Issuer:    "test",
		},
	}
	//生成token
	token, err := j.CreateToken(claims)
	if err != nil {
		response.Msg(c, http.StatusInternalServerError, 401, "token生成失败,重新再试", "")
		return ""
	}
	return token
}

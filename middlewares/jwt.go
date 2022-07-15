package middlewares

import (
	"dousheng/global"
	"dousheng/response"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	UserID   int
	UserName string
	Role     int
	jwt.StandardClaims
}

//错误类型
var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

//进行JWT验证函数
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//取token
		token := c.Query("token")
		if token == "" {
			response.Msg(c, http.StatusUnauthorized, 401, "请登录", "")
			c.Abort() //需先中断中间件，return不会中断中间件
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息，返回CustomClaims类型信息
		msg, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.Msg(c, http.StatusUnauthorized, 401, "授权已过期", "")
				c.Abort()
				return
			}
			response.Msg(c, http.StatusUnauthorized, 401, "未登陆", "")
			c.Abort()
			return
		}
		fmt.Println(c)
		// gin的上下文记录claims和userId的值
		c.Set("claims", msg)
		c.Set("userId", msg.UserID)
		c.Next()
	}
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.Settings.JWTKey.SigningKey),
	}
}

// 创建token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	//错误处理
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(30 * 24 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}

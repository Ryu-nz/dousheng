package controller

import (
	"dousheng/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

//注册请求
type RegisterReq struct {
	Username string `form:"username" json:"username" binding:"required,max=32"`
	Password string `form:"password" json:"password" binding:"required,max=32"`
}

//登录请求
type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//获取用户信息请求
type GetUserReq struct {
	UserID int    `form:"user_id" json:"user_id"`
	Token  string `form:"token" json:"token"`
}

//登录信息返回 userId + token
type LoginResp struct {
	Response
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

//用户信息返回
type UserResp struct {
	UserID        int    `form:"id" json:"id"`
	Username      string `form:"name" json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

//处理注册请求处
func UserRegister(c *gin.Context) {
	RegistrForm := RegisterReq{}
	//注意通过ShouldBind接收数据后再使用PostForm()无法再接到数据
	if err := c.ShouldBind(&RegistrForm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	user := User{Username: RegistrForm.Username, Password: RegistrForm.Password}

	result := global.DB.Select("Username", "Password").Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1})
		return
	}
	c.JSON(http.StatusOK, LoginResp{
		Response: Response{StatusCode: 0, StatusMsg: "用户注册成功"},
		UserID:   user.UserID,
		Token:    CreateToken(c, user.UserID, user.Username, user.Role),
	})
}

// 处理登录请求
func PasswordLogin(c *gin.Context) {
	LoginForm := LoginReq{}
	if err := c.ShouldBind(&LoginForm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	var user User
	global.DB.Where("username = ? and password = ?", LoginForm.Username, LoginForm.Password).Find(&user)
	if user.UserID != 0 {
		token := CreateToken(c, user.UserID, user.Username, user.Role)
		c.JSON(http.StatusOK, LoginResp{
			Response: Response{StatusCode: 0},
			UserID:   user.UserID,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1})
	}
}

// 处理获取自身信息请求
func GetUser(c *gin.Context) {
	GetUser := GetUserReq{}
	c.ShouldBind(&GetUser)
	Claims, err := NewJWT().ParseToken(GetUser.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error()})
		return
	} else if Claims.UserID != GetUser.UserID {
		c.JSON(http.StatusUnauthorized, Response{StatusCode: -1})
	} else {
		var user User
		global.DB.Where("user_id = ?", Claims.UserID).Find(&user)
		if user.UserID != 0 {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "success",
				"user":        GetUserResp(user, false),
			})
		}
	}
}

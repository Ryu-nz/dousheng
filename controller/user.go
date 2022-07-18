package controller

import (
	"dousheng/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

//处理注册请求
func UserRegister(c *gin.Context) {
	RegistrForm := RegisterReq{}
	//注意通过ShouldBind接收数据后再使用PostForm()无法再接到数据
	if err := c.ShouldBind(&RegistrForm); err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "请求数据错误"})
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
func UserLogin(c *gin.Context) {
	//接收数据
	LoginForm := LoginReq{}
	if err := c.ShouldBind(&LoginForm); err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: "请求数据出错"})
		return
	}
	//从数据库中查询数据
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
	//接收请求数据
	GetUser := UserReq{}
	c.ShouldBind(&GetUser)
	Claims, err := NewJWT().ParseToken(GetUser.Token)
	//解析token校验
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error()})
		return
	} else if Claims.UserID != GetUser.UserID {
		c.JSON(http.StatusUnauthorized, Response{StatusCode: -1})
	} else {
		var user User
		global.DB.Where("user_id = ?", Claims.UserID).Find(&user)
		//返回数据
		c.JSON(http.StatusOK, UserInfoResp{
			Response: Response{StatusCode: 0},
			UserResp: GetUserResp(user, false),
		})
	}
}

package service

import (
	"dousheng/global"
	"dousheng/middlewares"
	"dousheng/models"
	"dousheng/response"
	"dousheng/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

//register by username and password
func UserRegister(c *gin.Context) {
	RegistrForm := models.RegisterReq{}
	//注意通过ShouldBind接收数据后再使用PostForm()无法再接到数据
	if err := c.ShouldBind(&RegistrForm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	user := models.User{Username: RegistrForm.Username, Password: RegistrForm.Password}
	result := global.DB.Select("Username", "Password").Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  result.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "用户创建成功",
		"user_id":     user.UserID,
		"token":       utils.CreateToken(c, user.UserID, user.Username, user.Role),
	})
}

// login by username and password
func PasswordLogin(c *gin.Context) {
	LoginForm := models.LoginReq{}
	if err := c.ShouldBind(&LoginForm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	var user models.User
	global.DB.Where("username = ? and password = ?", LoginForm.Username, LoginForm.Password).Find(&user)
	if user.UserID != 0 {
		token := utils.CreateToken(c, user.UserID, user.Username, user.Role)
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"msg":     "登录成功",
			"user_id": user.UserID,
			"token":   token,
		})
	} else {
		response.Msg(c, 500, -1, "", "")
	}
}

//get the user's information by user_id and token
func GetUser(c *gin.Context) {
	GetUser := models.GetUser{}
	if err := c.ShouldBind(&GetUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	Claims, err := middlewares.NewJWT().ParseToken(GetUser.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	} else if Claims.UserID != GetUser.UserID {
		response.Msg(c, http.StatusInternalServerError, -1, "token和用户不匹配", "")
	} else {
		var user models.User
		global.DB.Where("user_id = ?", Claims.UserID).Find(&user)
		if user.UserID != 0 {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 0,
				"status_msg":  "success",
				"user":        models.GetUserResp(user, false),
			})
		}
	}
}

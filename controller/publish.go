package controller

import (
	"dousheng/global"
	"dousheng/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	//接收请求数据
	token := c.PostForm("token")
	title := c.PostForm("title")
	file, _ := c.FormFile("data")
	//解析token获取user_id
	Msg, err := NewJWT().ParseToken(token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: "token数据错误"})
	}
	user_id := Msg.UserID
	//根据vid设置存储文件名
	video := Video{}
	global.DB.Last(&video)
	fileVid := video.Vid + 1
	filename := fmt.Sprint(fileVid) + ".mp4"
	//获取文件内容
	fileObj, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: "文件获取失败"})
		return
	}
	//调用上传文件
	if ok, err := utils.UploadFile("video", filename, fileObj, file.Size); !ok {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "文件上传失败"})
		return
	}
	videoLatest := Video{
		UID:        user_id,
		Title:      title,
		PlayURL:    "http://" + utils.GetIP() + ":9001/video/" + filename,
		CreateTime: time.Now(), //创建时间
	}
	result := global.DB.Create(&videoLatest)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "插入数据库出错"})
		return
	}
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "视频上传成功"})
}

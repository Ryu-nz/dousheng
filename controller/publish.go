package controller

import (
	"dousheng/global"
	"dousheng/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//处理发布视频
func Publish(c *gin.Context) {
	//接收请求数据
	actionReq := PublishAction{}
	if err := c.ShouldBind(actionReq); err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "请求数据出错"})
	}
	//解析token获取user_id
	Msg, err := NewJWT().ParseToken(actionReq.Token)
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
	fileObj, err := actionReq.File.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: "文件获取失败"})
		return
	}
	//调用上传文件
	if ok, err := utils.UploadFile("video", filename, fileObj, actionReq.File.Size); !ok {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "文件上传失败"})
		return
	}
	videoLatest := Video{
		UID:        user_id,
		Title:      actionReq.Title,
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

//获取发布视频
func PublishList(c *gin.Context) {
	//接收请求数据并与token比对校验
	ListReq := ListReq{}
	if err := c.ShouldBind(&ListReq); err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "请求参数出错"})
		return
	}
	if Msg, err := NewJWT().ParseToken(ListReq.Token); err != nil || Msg.UserID != ListReq.UserID {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: "token数据错误"})
		return
	}
	//根据user_id获取user和video
	user, video, videoList := User{}, []Video{}, []VideoResp{}
	global.DB.Find(&user, ListReq.UserID)
	userResp := GetUserResp(user, false)
	global.DB.Where("uid = ?", ListReq.UserID).Find(&video)
	//构造video_list返回
	for _, v := range video {
		videoResp := GetVideoResp(v, userResp, false)
		videoList = append(videoList, videoResp)
	}
	c.JSON(http.StatusOK, ListResp{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
	})
}

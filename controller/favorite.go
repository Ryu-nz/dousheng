package controller

import (
	"dousheng/global"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	//接收数据
	actionReq := FavoriteActionReq{}
	if err := c.ShouldBind(&actionReq); err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "请求数据错误"})
		return
	}
	//解析token获取user_id
	Msg, err := NewJWT().ParseToken(actionReq.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: "token数据错误"})
		return
	}
	user_id := Msg.UserID
	//根据操作类型修改数据库
	favorite := Favorite{Uid: user_id, Vid: actionReq.VideoID}
	if actionReq.ActionType == 1 {
		favorite.IsFavorite = true
		if result := global.DB.Create(&favorite); result.Error != nil {
			c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: result.Error.Error() + "数据插入错误"})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else if actionReq.ActionType == 2 {
		if err := global.DB.Model(&Favorite{}).Where("uid = ? and vid = ?", user_id, actionReq.VideoID).
			Update("is_favorite", false).Error; err != nil {
			c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "数据修改错误"})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: "无对应操作类型"})
		return
	}
}

//与publish list代码相似度很大，是否可以合并
func FavoriteList(c *gin.Context) {
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
	user, favorite, videoList := User{}, []Favorite{}, []VideoResp{}
	global.DB.Find(&user, ListReq.UserID)
	userResp := GetUserResp(user, false)
	global.DB.Where("uid = ?", ListReq.UserID).Find(&favorite)
	//构造video_list返回
	for _, f := range favorite {
		video := Video{}
		global.DB.Find(&video, f.Vid)
		videoResp := GetVideoResp(video, userResp, false)
		videoList = append(videoList, videoResp)
	}
	c.JSON(http.StatusOK, ListResp{
		Response:  Response{StatusCode: 0},
		VideoList: videoList,
	})
}

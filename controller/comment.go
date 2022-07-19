package controller

import (
	"dousheng/global"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	//接收数据
	commentReq := CommentActionReq{}
	Msg, err := NewJWT().ParseToken(commentReq.Token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "Token解析错误"})
		return
	}
	if err := c.ShouldBind(&commentReq); err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "请求数据错误"})
		return
	}
	//操作数据库
	if commentReq.ActionType != 1 && commentReq.ActionType != 2 {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "操作类型错误"})
		return
	} else if commentReq.ActionType == 1 { //插入数据
		comment := Comment{
			Uid:         Msg.UserID,
			Vid:         commentReq.VideoID,
			CommentText: commentReq.CommentText,
			CreateTime:  time.Now()}
		result := global.DB.Create(&comment)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "创建数据错误"})
			return
		}
		c.JSON(http.StatusOK, CommentActionResp{
			Response:    Response{StatusCode: 0},
			CommentResp: GetCommentResp(comment),
		})
	} else { //删除数据
		result := global.DB.Delete(&Comment{}, commentReq.CommentId)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "删除数据错误"})
			return
		}
		c.JSON(http.StatusOK, CommentActionResp{
			Response:    Response{StatusCode: 0},
			CommentResp: CommentResp{},
		})
	}
}

func CommentList(c *gin.Context) {
	listReq := CommentListReq{}
	if err := c.ShouldBind(&listReq); err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: -1, StatusMsg: err.Error() + "请求数据错误"})
		return
	}

	comment := []Comment{}
	commentList := []CommentResp{}
	global.DB.Where("vid = ?", listReq.VideoID).Find(&comment)
	for _, v := range comment {
		commentList = append(commentList, GetCommentResp(v))
	}
	c.JSON(http.StatusOK, CommentListResp{
		Response:    Response{StatusCode: 0},
		CommentList: commentList,
	})
}

package controller

import (
	"dousheng/global"

	"github.com/gin-gonic/gin"
)

//通过Video数据和user数据获取返回视频数据
func GetVideoResp(Video Video, auther UserResp, IsFavorite bool) VideoResp {
	VideoResp := VideoResp{
		Vid:           Video.Vid,
		Auther:        auther,
		PlayURL:       Video.PlayURL,
		CoverURL:      Video.CoverURL,
		FavoriteCount: Video.FavoriteCount,
		CommentCount:  Video.CommentCount,
		IsFavorite:    IsFavorite,
		Title:         Video.Title,
	}
	return VideoResp
}

func Feed(c *gin.Context) {
	//接收请求数据
	FeedReq := FeedReq{}
	c.ShouldBind(&FeedReq)
	//从数据库中查找Video，结合auther构造VideoResp
	video, feed, user := []Video{}, []VideoResp{}, User{}
	max := 30
	global.DB.Find(&video)
	if len(video) > max {
		video = video[:max]
	}
	for _, v := range video {
		global.DB.Where("user_id = ?", v.Uid).Find(&user)
		//关注待实现
		userResp := GetUserResp(user, false)
		//点赞待实现
		videoResp := GetVideoResp(v, userResp, false)
		feed = append(feed, videoResp)
	}

	c.JSON(200, FeedResp{
		Response:  Response{StatusCode: 0},
		VideoList: feed,
	})
}

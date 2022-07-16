package controller

import (
	"dousheng/global"

	"github.com/gin-gonic/gin"
)

//视频Feed流请求
type FeedReq struct {
	LatestTime int    `form:"latest_time" json:"latest_time"`
	Token      string `form:"token" json:"token"`
}

//视频发布请求
type VideoPublishReq struct {
	Token string `form:"token" json:"token" binding:"required"`
	Data  []byte `form:"data" json:"data" binding:"required"`
	Title string `form:"titile" json:"title" binding:"required"`
}


//返回视频结构
type VideoResp struct {
	Vid           int      `form:"id" json:"id"`
	Auther        UserResp `form:"auther" json:"auther"`
	PlayURL       string   `json:"play_url"`
	CoverURL      string   ` json:"cover_url"`
	FavoriteCount int      `json:"favorite_count"`
	CommentCount  int      `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   ` json:"title"`
}

//返回视频流结构
type FeedResp struct {
	Response
	VideoList []VideoResp `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

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
	FeedReq := FeedReq{}
	c.ShouldBind(&FeedReq) //接收请求数据
	var Video []Video
	var feed []VideoResp
	var user User
	db := global.DB
	max := 30

	db.Find(&Video) //从数据中查找Video，构造VideoResp
	if len(Video) > max {
		Video = Video[:max]
	}
	for _, v := range Video {
		db.Where("user_id = ?", v.UID).Find(&user)
		userResp := GetUserResp(user, false)
		media := GetVideoResp(v, userResp, false)
		feed = append(feed, media)
	}

	c.JSON(200, FeedResp{
		Response: Response{StatusCode: 0},
		VideoList:     feed,
	})
}

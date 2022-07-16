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
type VedioPublishReq struct {
	Token string `form:"token" json:"token" binding:"required"`
	Data  []byte `form:"data" json:"data" binding:"required"`
	Title string `form:"titile" json:"title" binding:"required"`
}


//返回视频结构
type VedioResp struct {
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
	Feed []VedioResp
}

//通过vedio数据和user数据获取返回视频数据
func GetVedioResp(vedio Vedio, auther UserResp, IsFavorite bool) VedioResp {
	VedioResp := VedioResp{
		Vid:           vedio.Vid,
		Auther:        auther,
		PlayURL:       vedio.PlayURL,
		CoverURL:      vedio.CoverURL,
		FavoriteCount: vedio.FavoriteCount,
		CommentCount:  vedio.CommentCount,
		IsFavorite:    IsFavorite,
		Title:         vedio.Title,
	}
	return VedioResp
}

func Feed(c *gin.Context) {
	FeedReq := FeedReq{}
	c.ShouldBind(&FeedReq) //接收请求数据
	var vedio []Vedio
	var feed []VedioResp
	var user User
	db := global.DB
	max := 30

	db.Find(&vedio) //从数据中查找vedio，构造vedioResp
	if len(vedio) > max {
		vedio = vedio[:max]
	}
	for _, v := range vedio {
		db.Where("user_id = ?", v.UID).Find(&user)
		userResp := GetUserResp(user, false)
		media := GetVedioResp(v, userResp, false)
		feed = append(feed, media)
	}

	c.JSON(200, FeedResp{
		Response: Response{StatusCode: 0},
		Feed:     feed,
	})
}

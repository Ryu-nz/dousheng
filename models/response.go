package models

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

type UserResp struct {
	UserID        int    `form:"id" json:"id"`
	Username      string `form:"name" json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

func GetUserResp(user User, IsFollow bool) UserResp {
	UserResp := UserResp{
		UserID:        user.UserID,
		Username:      user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      IsFollow,
	}
	return UserResp
}

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
package controller

import "time"


type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}


//用户对应结构体
type User struct {
	UserID        int    `gorm:"column:user_id;primary_key" json:"user_id"`
	Username      string `gorm:"column:username" json:"username"`
	Password      string `gorm:"column:password" json:"password"`
	FollowCount   int    `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount int    `gorm:"column:follower_count" json:"follower_count"`
	Role          int    `gorm:"column:role" json:"role"`
	HeadUrl       string `gorm:"head_url" json:"head_url"`
}

//视频对应结构
type Vedio struct {
	Vid           int       `gorm:"column:vid;primary_key" json:"vid"`
	UID           int       `gorm:"column:uid" json:"uid"`
	Title         string    `gorm:"column:title" json:"title"`
	PlayURL       string    `gorm:"column:play_url" json:"play_url"`
	CoverURL      string    `gorm:"column:cover_url" json:"cover_url"`
	FavoriteCount int       `gorm:"column:favorite_count" json:"favorite_count"`
	CommentCount  int       `gorm:"column:comment_count" json:"comment_count"`
	CreateTime    time.Time `gorm:"cloumn:create_time" json:"create_time"`
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

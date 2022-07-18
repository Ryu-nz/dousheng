package controller

import "mime/multipart"

//user request
//注册请求
type RegisterReq struct {
	Username string `form:"username" json:"username" binding:"required,max=32"`
	Password string `form:"password" json:"password" binding:"required,max=32"`
}

//登录请求
type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

//获取用户信息请求
type UserReq struct {
	UserID int    `form:"user_id" json:"user_id"`
	Token  string `form:"token" json:"token"`
}

//feed request
//视频Feed流请求
type FeedReq struct {
	LatestTime int    `form:"latest_time" json:"latest_time"`
	Token      string `form:"token" json:"token"`
}

//publish request
//publish action请求
type PublishActionReq struct {
	Token string                `form:"token" binding:"required"`
	Title string                `form:"title" binding:"required"`
	File  *multipart.FileHeader `form:"data" binding:"required"`
}

//publish list请求
type ListReq struct {
	UserID int    `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

//favorite request
type FavoriteActionReq struct {
	Token      string `form:"token" binding:"required"`
	VideoID    int    `form:"video_id" binding:"required"`
	ActionType int    `form:"action_type" binding:"required"`
}

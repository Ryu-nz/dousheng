package controller

import "mime/multipart"

//user request
type RegisterReq struct {
	Username string `form:"username" json:"username" binding:"required,max=32"`
	Password string `form:"password" json:"password" binding:"required,max=32"`
}

type LoginReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type UserReq struct {
	UserID int    `form:"user_id" json:"user_id"`
	Token  string `form:"token" json:"token"`
}

//feed request
type FeedReq struct {
	LatestTime int    `form:"latest_time" json:"latest_time"`
	Token      string `form:"token" json:"token"`
}

//publish request
type PublishActionReq struct {
	Token string                `form:"token" binding:"required"`
	Title string                `form:"title" binding:"required"`
	File  *multipart.FileHeader `form:"data" binding:"required"`
}

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

//comment request
type CommentActionReq struct {
	Token       string `form:"token" binding:"required"`
	VideoID     int    `form:"video_id" binding:"required"`
	ActionType  int    `form:"action_type" binding:"required"`
	CommentText string `form:"comment_text"`
	CommentId   int    `form:"comment_id"`
}

type CommentListReq struct {
	Token   string `form:"token"`
	VideoID int    `form:"video_id" binding:"required"`
}

package controller

import "dousheng/global"

//user response
type LoginResp struct {
	Response
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfoResp struct {
	Response
	UserResp
}

type UserResp struct {
	UserID        int    `form:"id" json:"id"`
	Username      string `form:"name" json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

//从user对象获取用户返回结构
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

//feed、video response
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

type FeedResp struct {
	Response
	VideoList []VideoResp `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`
}

//publish response
type ListResp struct {
	Response
	VideoList []VideoResp `json:"video_list,omitempty"`
}

//comment response
type CommentActionResp struct {
	Response
	CommentResp
}

type CommentResp struct {
	ID         int      `json:"id"`
	User       UserResp `json:"user"`
	Content    string   `json:"content"`
	CreateDate string   `json:"create_date"`
}

func GetCommentResp(comment Comment) CommentResp {
	resp, user := CommentResp{}, User{}
	resp.ID = comment.CommentID
	global.DB.Find(&user, comment.Uid)
	if user.UserID == 0 {
		return CommentResp{}
	}
	resp.User = GetUserResp(user, false)
	resp.Content = comment.CommentText
	resp.CreateDate = comment.CreateTime.Format("01-02")
	return resp
}

type CommentListResp struct {
	Response
	CommentList []CommentResp `json:"comment_list"`
}

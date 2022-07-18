package controller

//user response
//登录信息返回 userId + token
type LoginResp struct {
	Response
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

//查询信息返回
type UserInfoResp struct {
	Response
	UserResp
}

//返回用户信息格式
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
	NextTime  int64       `json:"next_time,omitempty"`
}

//publish response

//publish list 返回
type ListResp struct {
	Response
	VideoList []VideoResp `json:"video_list,omitempty"`
}

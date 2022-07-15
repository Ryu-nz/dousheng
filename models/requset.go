package models

//定义接收用户传入参数表单并作基础限制

type RegisterReq struct {
	//用户名
	Username string `form:"username" json:"username" binding:"required,max=32"`
	//密码
	Password string `form:"password" json:"password" binding:"required,max=32"`
}

type LoginReq struct {
	//用户名
	Username string `form:"username" json:"username" binding:"required"`
	//密码
	Password string `form:"password" json:"password" binding:"required"`
}

type GetUser struct {
	//user_id
	UserID int `form:"user_id" json:"user_id" binding:"required"`
	//token
	Token string `form:"token" json:"token" binding:"required"`
}

//视频相关

type FeedReq struct {
	LatestTime int    `form:"latest_time" json:"latest_time"`
	Token      string `form:"token" json:"token"`
}

type VedioPublish struct {
	Token string `form:"token" json:"token" binding:"required"`
	Data  []byte `form:"data" json:"data" binding:"required"`
	Title string `form:"titile" json:"title" binding:"required"`
}

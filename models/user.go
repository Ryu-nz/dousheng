package models

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

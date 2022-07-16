package models

import "time"

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

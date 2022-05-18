package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Post    *Post
	PostId  uint `gorm:"not null"`
	User    *User
	UserId  uint `gorm:"not null"`
	Reply   *Comment
	ReplyId uint   `gorm:"not null"`
	Content string `gorm:"type:text;not null"`
}

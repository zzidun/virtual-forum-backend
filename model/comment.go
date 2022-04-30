package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Reply         *Comment
	ReplyId       uint `gorm:"not null"`
	User          *User
	UserId        uint `gorm:"not null"`
	CommentInfo   *CommentInfo
	CommentInfoId uint `gorm:"not null"`
}

type CommentInfo struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
}

package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Post    *Post
	PostId  uint `gorm:"not null"`
	User    *User
	UserId  uint `gorm:"not null"`
	Reply   *Comment
	ReplyId uint
	Content string `gorm:"type:text;not null"`
}

type CommentPostForm struct {
	PostId  string `json:"postid" binding:"required"`
	UserId  string `json:"userid" binding:"required"`
	ReplyId string `json:"replyid"`
	Content string `json:"content" binding:"required"`
}

package model

import "gorm.io/gorm"

// 实体

type Comment struct {
	gorm.Model
	PostId  uint `gorm:"not null"`
	User    *User
	UserId  uint `gorm:"not null"`
	ReplyId uint
	Content string `gorm:"type:text;not null"`
}

// 报文

type CommentPostForm struct {
	PostId  string `json:"postid" binding:"required"`
	UserId  string `json:"userid" binding:"required"`
	ReplyId string `json:"replyid"`
	Content string `json:"content" binding:"required"`
}

type CommentDeleteForm struct {
	CommentId  string `json:"commentid" binding:"required"`
	UserId  string `json:"userid" binding:"required"`
}

type CommentListEnrty struct {
	Id       string `json:"id" binding:"required"`
	UserName string `json:"username" binding:"required"`
	ReplyId  string `json:"replyId" binding:"required"`
	Content  string `json:"content" binding:"required"`
}

package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Category   *Category
	CategoryId uint   `gorm:"not null"`
	Title      string `gorm:"type:varchar(50);not null"`
	Speak      uint   `gorm:"not null"`
	User       *User
	UserId     uint `gorm:"not null"`
}

type PostPostForm struct {
	CategoryId string `json:"categoryid" binding:"required"`
	Title      string `json:"title" binding:"required"`
	UserId     string `json:"userid" binding:"required"`
}

type PostListRequired struct {
	CategoryId string `json:"categoryid" binding:"required"`
	Left       string `json:"left" binding:"required"`
	Right      string `json:"right" binding:"required"`
}

type PostListEntry struct {
	Id       string `json:"id" binding:"required"`
	Title    string `json:"name" binding:"required"`
	Speak    string `json:"speak" binding:"required"`
	UserName string `json:"username" binding:"required"`
}

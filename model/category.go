package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string `gorm:"type:varchar(20);not null;unique"`
	Speak  uint   `gorm:"not null"`
	Follow uint   `gorm:"not null"`
	Wiki   *Post
	WikiId uint
}

type Categoryer struct {
	gorm.Model
	Category   *Category
	CategoryId uint `gorm:"not null"`
	User       *User
	UserId     uint `gorm:"not null"`
	AdminType  bool `gorm:"type:bool;not null"`
}

type CategoryCreateForm struct {
	Name string `json:"name" binding:"required"`
}

type CategoryDeleteForm struct {
	CategoryId string `json:"categoryid" binding:"required"`
}

type CategoryListEnrty struct {
	Id         string `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Speak      string `json:"speak" binding:"required"`
	Follow     string `json:"follow" binding:"required"`
	Categoryer string `json:"categoryer" binding:"required"`
}

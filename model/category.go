package model

import "gorm.io/gorm"

// 实体

type Category struct {
	gorm.Model
	Name   string `gorm:"type:varchar(20);not null;unique;uniqueIndex"`
	Speak  uint   `gorm:"not null"`
	Follow uint   `gorm:"not null"`
	WikiId uint
}

type Categoryer struct {
	gorm.Model
	Category   *Category
	CategoryId uint `gorm:"not null"`
	UserId     uint `gorm:"not null"`
	Type       uint `gorm:"type:bool;not null"`
}

// 报文

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

type CategoryerForm struct {
	UserId string `json:"userid" binding:"required"`
	Type   string `json:"type" binding:"required"`
}

type CategoryWikiForm struct {
	CategoryId string `json:"categoryid" binding:"required"`
	PostId     string `json:"postid" binding:"required"`
}

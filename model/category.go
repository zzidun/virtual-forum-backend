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

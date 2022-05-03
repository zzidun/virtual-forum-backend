package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name   string `gorm:"type:varchar(20);not null;unique"`
	speak  uint   `gorm:"not null"`
	follow uint   `gorm:"not null"`
	wiki   *Post
	wikiId uint `gorm:"not null"`
}

type CategoryAdmin struct {
	gorm.Model
	Category   *Category
	CategoryId uint `gorm:"not null"`
	User       *User
	UserId     uint `gorm:"not null"`
	AdminType  bool `gorm:"type:bool;not null"`
}

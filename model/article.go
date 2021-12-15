package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	User_ID    uint `gorm:"not null"`
	CategoryID uint `gorm:"not null"`
	Category   *Category
	Title      string `gorm:"type:varchar(50);not null"`
	Content    string `gorm:"text;not null"`
}

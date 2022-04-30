package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Category   *Category
	CategoryId uint   `gorm:"not null"`
	Title      string `gorm:"type:varchar(50);not null"`
}

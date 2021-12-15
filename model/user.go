package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(50);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

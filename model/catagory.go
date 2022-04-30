package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;unique"`
}

type CategoryAdmin struct {
	gorm.Model
}

type CategoryUser struct {
	gorm.Model
	
}
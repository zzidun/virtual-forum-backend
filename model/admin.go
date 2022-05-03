package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name         string `gorm:"type:varchar(20);not null;unique"`
	Password     string `gorm:"size:255;not null"`
	AdminGroup   *AdminGroup
	AdminGroupId uint `gorm:"not null"`
}

type AdminGroup struct {
	gorm.Model
	Name         string `gorm:"type:varchar(20);not null;unique"`
	AdminPerm    bool   `gorm:"type:bool;not null"`
	BanPerm      bool   `gorm:"type:bool;not null"`
	CategoryPerm bool   `gorm:"type:bool;not null"`
	PostPerm     bool   `gorm:"type:bool;not null"`
}

type AdminLoginForm struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

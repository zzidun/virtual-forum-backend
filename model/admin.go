package model

import "gorm.io/gorm"

//实体

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null;unique;uniqueIndex"`
	Password string `gorm:"size:255;not null"`
	Group    *AdminGroup
	GroupId  uint `gorm:"not null"`
}

type AdminGroup struct {
	gorm.Model
	Name         string `gorm:"type:varchar(20);not null;unique;uniqueIndex"`
	AdminPerm    bool   `gorm:"type:bool;not null"`
	BanPerm      bool   `gorm:"type:bool;not null"`
	CategoryPerm bool   `gorm:"type:bool;not null"`
	PostPerm     bool   `gorm:"type:bool;not null"`
}

// 报文

type AdminLoginForm struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type AdminCreateForm struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	GroupId  string `json:"groupid" binding:"required"`
}

type AdminGroupCreateForm struct {
	Name         string `json:"name" binding:"required"`
	AdminPerm    bool   `json:"adminperm" binding:"required"`
	BanPerm      bool   `json:"banperm" binding:"required"`
	CategoryPerm bool   `json:"categoryperm" binding:"required"`
	PostPerm     bool   `json:"postperm" binding:"required"`
}

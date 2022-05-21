package model

import "gorm.io/gorm"

//实体

type Admin struct {
	gorm.Model
	AdminPerm    uint `gorm:"not null"`
	BanPerm      uint `gorm:"not null"`
	CategoryPerm uint `gorm:"not null"`
}

// 报文

type AdminForm struct {
	UserId       string `json:"userid" binding:"required"`
	AdminPerm    string `json:"adminperm" binding:"required"`
	BanPerm      string `json:"banperm" binding:"required"`
	CategoryPerm string `json:"categoryperm" binding:"required"`
}

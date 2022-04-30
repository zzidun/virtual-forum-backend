package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Email    string `gorm:"type:varchar(50);not null;unique"`
	Password string `gorm:"size:255;not null"`
}

type AdminGroup struct {
	gorm.Model
	Name             string `gorm:"type:varchar(20);not null;unique"`
	AdminGroupPerm   *AdminGroupPerm
	AdminGroupPermId uint `gorm:"not null"`
}

type AdminGroupMember struct {
	gorm.Model
	Admin        *Admin
	AdminId      uint `gorm:"not null"`
	AdminGroup   *AdminGroup
	AdminGroupId uint `gorm:"not null"`
}

type AdminGroupPerm struct {
	gorm.Model
	UpdateAdminGroup bool `gorm:"type:bool;not null"`
	UpdateAdmin      bool `gorm:"type:bool;not null"`
	UpdateBanIp      bool `gorm:"type:bool;not null"`
	UpdateBanUser    bool `gorm:"type:bool;not null"`
	UpdateForum      bool `gorm:"type:bool;not null"`
	UpdatePost       bool `gorm:"type:bool;not null"`
}

// type AdminOperator struct {
// 	gorm.Model

// }

type AdminLoginForm struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"name" binding:"required"`
}

package model

import "gorm.io/gorm"

// 实体

type BannedIpv4 struct {
	gorm.Model
	LeftIp  string `gorm:"type:varchar(32);not null;unique"`
	RightIp string `gorm:"type:varchar(32);not null;unique"`
	UserId  uint   `gorm:"not null"`
}

type FailedUser struct {
	gorm.Model
	FailedIp     string `gorm:"type:varchar(32);not null;unique"`
	FailedUserId uint   `gorm:"not null"`
	Trys         uint   `gorm:"not null"`
}

// 报文

type BannedIpv4Form struct {
	LeftIp  string `gorm:"type:varchar(32);not null;unique"`
	RightIp string `gorm:"type:varchar(32);not null;unique"`
	UserId  string `gorm:"not null"`
}

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"type:varchar(20);not null"`
	Email       string `gorm:"type:varchar(50);not null;unique"`
	Password    string `gorm:"size:255;not null"`
	Emailverify bool   `gorm:"type:bool;not null"`
	UserInfo    *UserInfo
	UserInfoId  uint `gorm:"not null"`
}

type UserInfo struct {
	gorm.Model
	Signal string `gorm:"size:255;not null"`
}

// type UserOperator struct {
// 	gorm.Model
// }

type UserBanned struct {
	gorm.Model
}

type UserShield struct {
	gorm.Model
	User         *User
	UserId       uint `gorm:"not null"`
	ShieldUser   *User
	ShieldUserId uint `gorm:"not null"`
}

type UserCollect struct {
	gorm.Model
	User   *User
	UserId uint `gorm:"not null"`
	Post   *Post
	PostId uint `gorm:"not null"`
}

type UserFollow struct {
	gorm.Model
	User       *User
	UserId     uint `gorm:"not null"`
	Category   *Category
	CategoryId uint `gorm:"not null"`
}

type UserForm struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binging:"required"`
}

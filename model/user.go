package model

import "gorm.io/gorm"

// 报文

type User struct {
	gorm.Model
	Name          string `gorm:"type:varchar(20);not null;unique;uniqueIndex"`
	Email         string `gorm:"type:varchar(50);not null;unique;uniqueIndex"`
	Password      string `gorm:"size:255;not null"`
	Signal        string `gorm:"size:255"`
	LastLoginIpv4 string `gorm:"type:varchar(32);"`
	Speak         uint   `gorm:"not null"`
}

type UserFollow struct {
	gorm.Model
	User       *User
	UserId     uint `gorm:"not null"`
	Category   *Category
	CategoryId uint `gorm:"not null"`
	Count      uint `gorm:"not null"`
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

// 报文

type UserRegisterForm struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binging:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginForm struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateForm struct {
	UserId   string `json:"userid" binding:"required"`
	Email    string `json:"email" binging:"required"`
	Password string `json:"password" binding:"required"`
	Signal   string `json:"signal" binding:"required"`
}

type UserShieldForm struct {
	UserId      string `json:"userid1" binding:"required"`
	ShielUserId string `json:"userid2" binding:"required"`
}

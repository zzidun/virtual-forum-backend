package util

import (
	"gorm.io/gorm"
	"zzidun.tech/vforum0/model"
)

func Email_Exist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

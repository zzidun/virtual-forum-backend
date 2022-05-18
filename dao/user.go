package dao

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"zzidun.tech/vforum0/model"
)

// 创建用户帐号，数据验证，加密密码
func UserCreate(urform *model.UserRegisterForm) (err error) {

	password, err := bcrypt.GenerateFromPassword([]byte(urform.Password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("gen password failed", zap.Error(err))
		err = ErrorPasswordWrong
		return
	}

	// 组装用户对象
	user := model.User{
		Name:     urform.Name,
		Email:    urform.Email,
		Password: string(password),
	}

	db := DatabaseGet()
	if err = db.Create(&user).Error; err != nil {
		db.Rollback()
		zap.L().Error("insert user failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

func UserLogin(ulform *model.UserLoginForm) (id uint, err error) {

	var user model.User

	db := DatabaseGet()
	count := db.Where("name = ?", ulform.Name).Find(&user)

	if count.Error != nil {
		zap.L().Error("query user failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		err = ErrorNotExistFailed
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ulform.Password))
	if err != nil {
		zap.L().Error("compare password failed", zap.Error(err))
		err = ErrorPasswordWrong
		return
	}

	id = user.ID
	return
}

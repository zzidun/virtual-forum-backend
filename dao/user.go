package dao

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"zzidun.tech/vforum0/model"
)

func UserEmailReapetCheck(email string) (err error) {
	db := DatabaseGet()
	count := db.Where("email = ?", email).Find(&model.User{})
	if count.Error != nil {
		zap.L().Error("query user failed", zap.Error(count.Error))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected != 0 {
		err = ErrorExistFailed
		return
	}
	return
}

func UserNameReapetCheck(name string) (err error) {
	db := DatabaseGet()
	count := db.Where("name = ?", name).Find(&model.User{})

	if count.Error != nil {
		zap.L().Error("query user failed", zap.Error(count.Error))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected != 0 {
		err = ErrorExistFailed
		return
	}
	return
}

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

	if err = UserNameReapetCheck(user.Name); err != nil {
		zap.L().Error("insert user failed", zap.Error(err))
		return
	}

	if err = UserEmailReapetCheck(user.Email); err != nil {
		zap.L().Error("insert user failed", zap.Error(err))
		return
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

func UserUpdate(userId uint, email string, password string, signal string) (err error) {
	db := DatabaseGet()

	var user model.User
	count := db.Where("id = ?", userId).Find(&user)

	if count.Error != nil {
		zap.L().Error("query user failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		err = ErrorNotExistFailed
		return
	}

	if err = UserEmailReapetCheck(email); err != nil {
		zap.L().Error("insert user failed", zap.Error(err))
		return
	}

	user.Email = email
	user.Password = password
	user.Signal = signal

	if err = db.Save(&user).Error; err != nil {
		db.Rollback()
		zap.L().Error("update post speak count failed", zap.Error(err))
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
		zap.L().Error("query user failed", zap.Error(count.Error))
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

func UserQueryById(userId uint) (user model.User, err error) {

	db := DatabaseGet()
	count := db.Where("id = ?", userId).Find(&user)

	if count.Error != nil {
		zap.L().Error("query user failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		err = ErrorNotExistFailed
		return
	}

	return user, nil
}

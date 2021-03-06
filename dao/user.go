package dao

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"zzidun.tech/vforum0/model"
)

func UserEmailReapetCheck(email string) (err error) {
	db := DatabaseGet()
	count := db.Where("email = ?", email).Find(&model.User{})
	if count.Error != nil {
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
func UserCreate(name string, email string, passwordOrigin string) (user *model.User, err error) {

	password, err := bcrypt.GenerateFromPassword([]byte(passwordOrigin), bcrypt.DefaultCost)
	if err != nil {
		err = ErrorPasswordWrong
		return
	}

	if err = UserNameReapetCheck(name); err != nil {
		err = ErrorExistFailed
		return
	}

	if err = UserEmailReapetCheck(email); err != nil {
		err = ErrorExistFailed
		return
	}

	user = &model.User{
		Name:     name,
		Email:    string(email),
		Password: string(password),
		Speak:    0,
	}

	db := DatabaseGet()
	if err = db.Create(&user).Error; err != nil {
		db.Rollback()
		err = ErrorInsertFailed
		return
	}

	return
}

// 创建用户帐号，数据验证
func UserCreateWithBcrypted(name string, email string, password string) (user *model.User, err error) {

	if err = UserNameReapetCheck(name); err != nil {
		err = ErrorExistFailed
		return
	}

	if err = UserEmailReapetCheck(string(email)); err != nil {
		err = ErrorExistFailed
		return
	}

	user = &model.User{
		Name:     name,
		Email:    string(email),
		Password: string(password),
		Speak:    0,
	}

	db := DatabaseGet()
	if err = db.Create(&user).Error; err != nil {
		db.Rollback()
		err = ErrorInsertFailed
		return
	}

	return
}

func UserUpdate(userId uint, email string, passwordOrigin string, signal string, 
	passwordOriginOld string) (user *model.User, err error) {

	db := DatabaseGet()

	user, err = UserQueryById(userId)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwordOriginOld))
	if err != nil {
		err = ErrorPasswordWrong
		return
	}

	fmt.Println(user)

	password, err := bcrypt.GenerateFromPassword([]byte(passwordOrigin), bcrypt.DefaultCost)
	if err != nil {
		err = ErrorPasswordWrong
		return
	}

	if err = UserEmailReapetCheck(string(email)); err != nil {
		err = ErrorExistFailed
		return
	}

	user.Email = string(email)
	user.Password = string(password)
	user.Signal = signal

	if err = db.Save(&user).Error; err != nil {
		db.Rollback()
		err = ErrorInsertFailed
		return
	}

	return
}

func UserUpdateWithBcrypted(userId uint, email string, password string, signal string,
	emailOld string, passwordOld string) (user *model.User, err error) {

	db := DatabaseGet()

	user, err = UserQueryById(userId)
	if err != nil {
		return
	}

	if emailOld != user.Email || passwordOld != user.Password {
		err = ErrorPasswordWrong
		return
	}

	fmt.Println(user)

	if err = UserEmailReapetCheck(string(email)); err != nil {
		err = ErrorExistFailed
		return
	}

	user.Email = string(email)
	user.Password = string(password)
	user.Signal = signal

	if err = db.Save(&user).Error; err != nil {
		db.Rollback()
		err = ErrorInsertFailed
		return
	}

	return
}

func UserLogin(email string, password string) (user *model.User, err error) {

	db := DatabaseGet()
	count := db.Where("email = ?", email).Find(&user)

	fmt.Println(user)

	if count.Error != nil {
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		err = ErrorNotExistFailed
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		err = ErrorPasswordWrong
		return
	}

	return
}

func UserLoginWithBcrypted(name string, password string) (user *model.User, err error) {

	db := DatabaseGet()
	count := db.Where("name = ?", name).Find(&user)

	fmt.Println(user)

	if count.Error != nil {
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		err = ErrorNotExistFailed
		return
	}

	if user.Password != password {
		err = ErrorPasswordWrong
		return
	}

	return
}

func UserQueryById(userId uint) (user *model.User, err error) {

	db := DatabaseGet()
	count := db.Where("id = ?", userId).Find(&user)

	if count.Error != nil {
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		err = ErrorNotExistFailed
		return
	}

	return user, nil
}

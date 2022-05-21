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
func UserCreate(name string, emailOrigin string, passwordOrigin string) (user *model.User, err error) {

	password, err := bcrypt.GenerateFromPassword([]byte(passwordOrigin), bcrypt.DefaultCost)
	if err != nil {
		err = ErrorPasswordWrong
		return
	}

	if err = UserNameReapetCheck(name); err != nil {
		err = ErrorExistFailed
		return
	}

	email, err := bcrypt.GenerateFromPassword([]byte(emailOrigin), bcrypt.DefaultCost)
	if err != nil {
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

func UserUpdate(userId uint, emailOrigin string, passwordOrigin string, signal string) (user *model.User, err error) {
	db := DatabaseGet()

	user, err = UserQueryById(userId)
	if err != nil {
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(passwordOrigin), bcrypt.DefaultCost)
	if err != nil {
		err = ErrorPasswordWrong
		return
	}

	email, err := bcrypt.GenerateFromPassword([]byte(emailOrigin), bcrypt.DefaultCost)
	if err != nil {
		err = ErrorExistFailed
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

func UserLogin(name string, password string) (user *model.User, err error) {

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

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
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

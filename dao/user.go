package dao

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"zzidun.tech/vforum0/model"
)

// 创建用户帐号，数据验证，加密密码
func UserCreate(urform *model.UserRegisterForm) (err error) {

	password, err := bcrypt.GenerateFromPassword([]byte(urform.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 组装用户对象
	user := model.User{
		Name:     urform.Name,
		Email:    urform.Email,
		Password: string(password),
	}

	db := DatabaseGet()
	db.Create(user)
	return
}

func UserLogin(ulform *model.UserLoginForm) (id uint, err error) {

	var user model.User

	db := DatabaseGet()
	db.Where("name = ?", ulform.Name).First(&user)
	if user.ID == 0 {
		return 0, errors.New("密码错误")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ulform.Password))
	if err != nil {
		return 0, errors.New("密码错误")
	}

	return user.ID, nil
}

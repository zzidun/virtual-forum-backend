package dao

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"zzidun.tech/vforum0/model"
)

func AdminLogin(alform *model.AdminLoginForm) (id uint, err error) {

	var user model.Admin

	db := DatabaseGet()
	db.Where("name = ?", alform.Name).First(&user)
	if user.ID == 0 {
		return 0, errors.New("密码错误")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(alform.Password))
	if err != nil {
		return 0, errors.New("密码错误")
	}

	return user.ID, nil
}

// 创建管理员，并且检查是否重复，加密密码
func AdminCreate(alform *model.AdminLoginForm) (err error) {

	return
}

func AdminDelete(admin *model.Admin) (err error) {
	return
}

func AdminChange(admin *model.Admin) (err error) {
	return
}

func AdminQuery(admin *model.Admin) (err error) {
	return
}

func AdminQueryById(admin *model.Admin) (err error) {
	return
}

func AdmingroupCreate(admingroup *model.Admin) (err error) {
	db := DatabaseGet()
	db.Create(admingroup)
	return
}

func AdmingroupDelete(admin *model.Admin) (err error) {
	return
}

func AdmingroupChange(admin *model.Admin) (err error) {
	return
}

func AdmingroupQuery(admin *model.Admin) (err error) {
	return
}

func AdmingroupQueryById(admin *model.Admin) (err error) {
	return
}

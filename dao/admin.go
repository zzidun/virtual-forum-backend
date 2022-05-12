package dao

import (
	"golang.org/x/crypto/bcrypt"
	"zzidun.tech/vforum0/model"
)

func AdminLogin(admin *model.Admin) (err error) {
	return
}

// 创建管理员，并且检查是否重复，加密密码
func AdminCreate(admin *model.Admin) (err error) {

	password, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	admin.Password = string(password)

	db := DatabaseGet()
	db.Create(admin)
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

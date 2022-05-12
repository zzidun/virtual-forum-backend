package dao

import "zzidun.tech/vforum0/model"

func AdminLogin(admin *model.Admin) (err error) {
	return
}

func AdminCreate(admin *model.Admin) (err error) {
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

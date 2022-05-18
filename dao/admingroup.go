package dao

import "zzidun.tech/vforum0/model"

func AdmingroupCreate(agcform *model.AdminGroupCreateForm) (err error) {

	// 组装用户对象
	admingroup := model.AdminGroup{
		Name:         agcform.Name,
		AdminPerm:    agcform.AdminPerm,
		BanPerm:      agcform.BanPerm,
		CategoryPerm: agcform.CategoryPerm,
		PostPerm:     agcform.PostPerm,
	}

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

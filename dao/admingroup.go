package dao

import (
	"go.uber.org/zap"
	"zzidun.tech/vforum0/model"
)

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

func AdmingroupQueryById(admingroupId uint) (admingroup model.AdminGroup, err error) {

	db := DatabaseGet()

	count := db.Where("id = ?", admingroupId).Find(&admingroup)

	if count.Error != nil {
		zap.L().Error("query admingroup failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query admingroup failed", zap.Error(err))
		err = ErrorExistFailed
		return
	}

	return
}

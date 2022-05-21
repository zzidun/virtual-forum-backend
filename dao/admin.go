package dao

import (
	"zzidun.tech/vforum0/model"
)

// 创建管理员
func AdminCreate(userId uint) (admin *model.Admin, err error) {

	// 组装用户对象
	admin = &model.Admin{
		UserId:       userId,
		AdminPerm:    0,
		BanPerm:      0,
		CategoryPerm: 0,
	}

	db := DatabaseGet()
	if db.Create(&admin).Error != nil {
		err = ErrorInsertFailed
		return
	}

	user, err := UserQueryById(userId)
	if err != nil {
		return
	}

	user.AdminId = admin.ID

	if db.Save(user).Error != nil {
		err = ErrorUpdateFailed
		return
	}

	return
}

// 取消管理员权限
func AdminDelete(userId uint) (err error) {

	db := DatabaseGet()

	user, err := UserQueryById(userId)
	if err != nil {
		return
	}

	admin, err := AdminQueryById(user.AdminId)
	if err != nil {
		return
	}

	user.AdminId = 0

	if db.Save(user).Error != nil {
		return ErrorUpdateFailed
	}

	if err = db.Delete(&admin).Error; err != nil {
		err = ErrorDeleteFailed
		return
	}

	return
}

// 更新用户权限
func AdminUpdate(adminId uint, adminPerm uint, banPerm uint, categoryPerm uint) (err error) {
	db := DatabaseGet()

	admin, err := AdminQueryById(adminId)
	if err != nil {
		return
	}

	admin.AdminPerm = adminPerm
	admin.BanPerm = banPerm
	admin.CategoryPerm = categoryPerm

	if db.Save(&admin).Error != nil {
		return ErrorUpdateFailed
	}

	return
}

func AdminQuery(left int, right int) (adminList []model.Admin, totNum int64, curNum int64, err error) {
	db := DatabaseGet()

	count := db.Limit(right - left).Offset(left).Find(&adminList)

	if count.Error != nil {
		err = ErrorQueryFailed
	}
	curNum = count.RowsAffected
	db.Model(&model.Admin{}).Count(&totNum)

	return
}

func AdminQueryById(adminId uint) (admin *model.Admin, err error) {
	db := DatabaseGet()

	count := db.Where("id = ?", adminId).Find(&admin)

	if count.Error != nil {
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		err = ErrorNotExistFailed
		return
	}

	return
}

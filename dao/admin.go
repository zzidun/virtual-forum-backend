package dao

import (
	"errors"
	"strconv"

	"go.uber.org/zap"
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

func AdminReapetCheck(name string) (err error) {
	db := DatabaseGet()
	count := db.Where("name = ?", name).Find(&model.Admin{})

	if count.Error != nil {
		zap.L().Error("query admin failed", zap.Error(count.Error))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected != 0 {
		err = ErrorExistFailed
		return
	}

	return
}

// 创建管理员，并且检查是否重复，加密密码
func AdminCreate(acform *model.AdminCreateForm) (err error) {

	groupid, err := strconv.ParseUint(acform.GroupId, 10, 32)

	password, err := bcrypt.GenerateFromPassword([]byte(acform.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 组装用户对象
	admin := model.Admin{
		Name:     acform.Name,
		Password: string(password),
		GroupId:  uint(groupid),
	}

	if err = AdminReapetCheck(admin.Name); err != nil {
		zap.L().Error("insert user failed", zap.Error(err))
		return
	}

	db := DatabaseGet()
	db.Create(admin)
	return

}

func AdminDelete(adminId uint) (err error) {

	db := DatabaseGet()

	var admin model.Admin

	count := db.Where("id = ?", adminId).Find(&admin)

	if count.Error != nil {
		zap.L().Error("query admin failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query admin failed", zap.Error(err))
		err = ErrorNotExistFailed
		return
	}

	if err = db.Delete(&admin).Error; err != nil {
		zap.L().Error("delelte admin failed", zap.Error(err))
		err = ErrorDeleteFailed
		return
	}

	return
}

func AdminChangeGroup(adminId uint, groupId uint) (err error) {
	db := DatabaseGet()

	var admin model.Admin

	count := db.Where("id = ?", adminId).Find(&admin)

	if count.Error != nil {
		zap.L().Error("query admin failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query admin failed", zap.Error(err))
		err = ErrorNotExistFailed
		return
	}

	if err = db.Model(&admin).Update("group_id", groupId).Error; err != nil {
		db.Rollback()
		zap.L().Error("update userFollow failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

func AdminChangePassword(adminId uint, oldPassword string, newPassword string) (err error) {
	return
}

func AdminQuery(admin *model.Admin) (err error) {
	return
}

func AdminQueryById(adminId uint) (admin model.Admin, err error) {
	db := DatabaseGet()

	count := db.Where("id = ?", adminId).Find(&admin)

	if count.Error != nil {
		zap.L().Error("query category failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query category failed", zap.Error(err))
		err = ErrorNotExistFailed
		return
	}

	return
}

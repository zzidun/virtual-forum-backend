package dao

import (
	"go.uber.org/zap"
	"zzidun.tech/vforum0/model"
)

func UserShieldCreate(userId uint, shieldUserId uint) (err error) {

	db := DatabaseGet()

	shield := model.UserShield{
		UserId:       userId,
		ShieldUserId: shieldUserId,
	}
	if err = db.Create(&shield).Error; err != nil {
		db.Rollback()
		zap.L().Error("insert userFollow failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}
	return
}

func UserShieldDelete(shieldId uint) (err error) {
	db := DatabaseGet()

	var shield model.UserShield

	count := db.Where("id = ?", shieldId).Find(&shield)

	if count.Error != nil {
		zap.L().Error("query post failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query post failed", zap.Error(err))
		err = ErrorNotExistFailed
		return
	}

	if err = db.Delete(&shield).Error; err != nil {
		zap.L().Error("delelte shield failed", zap.Error(err))
		err = ErrorDeleteFailed
		return
	}

	return
}

func UserShieldQuery(userId uint, shieldUserId uint) (shield model.UserShield, err error) {

	db := DatabaseGet()
	count := db.Where("user_id = ? AND shield_user_id = ?", userId, shieldUserId).Find(&shield)

	if count.Error != nil {
		zap.L().Error("query UserShield failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}

	return
}

func UserShieldQueryByUser1(userId uint) (shield []model.UserShield, err error) {

	db := DatabaseGet()
	count := db.Where("user_id = ?", userId).Find(&shield)

	if count.Error != nil {
		zap.L().Error("query UserShield failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}

	return
}

func UserShieldQueryByUser2(userId uint) (shield []model.UserShield, err error) {

	db := DatabaseGet()
	count := db.Where("shield_user_id = ?", userId).Find(&shield)

	if count.Error != nil {
		zap.L().Error("query UserShield failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}

	return
}

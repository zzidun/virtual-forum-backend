package dao

import (
	"go.uber.org/zap"
	"zzidun.tech/vforum0/model"
)

func UserFollowCreate(userId uint, categoryId uint) (err error) {

	var follow model.UserFollow

	db := DatabaseGet()
	count := db.Unscoped().Where("user_id = ? AND category_id = ?", userId, categoryId).Find(&follow)

	if count.Error != nil {
		zap.L().Error("query userFollow failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		follow = model.UserFollow{
			UserId:     userId,
			CategoryId: categoryId,
			Count:      0,
		}
		if err = db.Create(&follow).Error; err != nil {
			db.Rollback()
			zap.L().Error("insert userFollow failed", zap.Error(err))
			err = ErrorInsertFailed
			return
		}
	} else {
		if err = db.Unscoped().Model(&follow).Update("deleted_at", nil).Error; err != nil {
			db.Rollback()
			zap.L().Error("update userFollow failed", zap.Error(err))
			err = ErrorInsertFailed
			return
		}
	}
	return
}

func UserFollowDelete(userFollowId uint) (err error) {

	return
}

func UserFollowQuery(userId uint, categoryId uint) (userFollow model.UserFollow, err error) {
	var follow model.UserFollow

	db := DatabaseGet()
	count := db.Where("user_id = ? AND category_id = ?", userId, categoryId).Find(&follow)

	if count.Error != nil {
		zap.L().Error("query userFollow failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}

	return
}

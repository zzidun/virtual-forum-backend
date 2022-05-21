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

	if err = CategoryCountFollow(uint(categoryId), 1); err != nil {
		return
	}


	return
}

func UserFollowDelete(followId uint) (err error) {
	db := DatabaseGet()

	var follow model.UserFollow

	count := db.Where("id = ?", followId).Find(&follow)

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

	categoryId := follow.CategoryId
	if err = db.Delete(&follow).Error; err != nil {
		zap.L().Error("delelte follow failed", zap.Error(err))
		err = ErrorDeleteFailed
		return
	}

	if err = CategoryCountFollow(uint(categoryId), 0); err != nil {
		return
	}

	return
}

func UserFollowQuery(userId uint, categoryId uint) (follow model.UserFollow, err error) {

	db := DatabaseGet()
	count := db.Where("user_id = ? AND category_id = ?", userId, categoryId).Find(&follow)

	if count.Error != nil {
		zap.L().Error("query userFollow failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}

	return
}

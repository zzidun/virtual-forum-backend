package dao

import (
	"go.uber.org/zap"
	"zzidun.tech/vforum0/model"
)

func CategoryerSet(categoryId uint, userId uint, adminType bool) (err error) {
	categoryer := model.Categoryer{
		CategoryId: categoryId,
		UserId:     userId,
		AdminType:  adminType,
	}

	db := DatabaseGet()
	if err = db.Create(&categoryer).Error; err != nil {
		db.Rollback()
		zap.L().Error("insert categoryer failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

func CategoryerCancel(categoryerId uint) (err error) {

	db := DatabaseGet()

	var categoryer model.Categoryer

	count := db.Where("id = ?", categoryerId).Find(&categoryer)

	if count.Error != nil {
		zap.L().Error("query categoryer failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query categoryer failed", zap.Error(err))
		err = ErrorNotExistFailed
		return
	}

	if err = db.Delete(&categoryer).Error; err != nil {
		zap.L().Error("delelte category failed", zap.Error(err))
		err = ErrorDeleteFailed
		return
	}

	return
}

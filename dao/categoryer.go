package dao

import (
	"go.uber.org/zap"
	"zzidun.tech/vforum0/model"
)

func CategoryerSet(categoryId uint, userId uint, categoryType uint) (err error) {
	categoryer := model.Categoryer{
		CategoryId: categoryId,
		UserId:     userId,
		Type:       categoryType,
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

// 查询版块的大版主,返回id
func CategoryerQueryByCategoryId(categoryId uint) (userId uint, err error) {

	db := DatabaseGet()

	var categoryer model.Categoryer

	count := db.Where("id = ? AND type = ?", categoryId, 1).Find(&categoryer)

	if count.Error != nil {
		zap.L().Error("query categoryer failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		userId = 0
		return
	}

	userId = categoryer.UserId

	return
}

package dao

import (
	"go.uber.org/zap"
	"zzidun.tech/vforum0/model"
)

func CategoryCreate(categoryName string) (err error) {
	category := model.Category{
		Name:   categoryName,
		Speak:  0,
		Follow: 0,
	}

	db := DatabaseGet()

	if err = db.Create(&category).Error; err != nil {
		db.Rollback()
		zap.L().Error("insert category failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

func CategoryerUpdate(categoryId uint, userId uint) (err error) {
	categoryer := model.Categoryer{
		CategoryId: categoryId,
		UserId:     userId,
		AdminType:  true,
	}

	db := DatabaseGet()
	db.Create(categoryer)

	return
}

func CategoryDelete(categoryId uint) (err error) {

	db := DatabaseGet()

	var category model.Category

	count := db.Where("id = ?", categoryId).Find(&category)

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

	if err = db.Delete(&category).Error; err != nil {
		zap.L().Error("delelte category failed", zap.Error(err))
		err = ErrorDeleteFailed
		return
	}

	return
}

func CategoryQuery() (category []model.Category, err error) {
	db := DatabaseGet()
	
	if db.Find(&category).Error != nil {
		err = ErrorQueryFailed
	}
	return
}

func CategoryQueryById(categoryId uint) (category model.Category, err error) {

	return
}

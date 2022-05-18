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

func CategoryQuery() {

}

func CategoryQueryById(categoryId uint) (category model.Category, err error) {
	return
}

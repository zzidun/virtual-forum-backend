package dao

import (
	"go.uber.org/zap"
	"zzidun.tech/vforum0/model"
)

func CategoryReapetCheck(name string) (err error) {
	db := DatabaseGet()
	count := db.Where("name = ?", name).Find(&model.Category{})

	if count.Error != nil {
		zap.L().Error("query category failed", zap.Error(count.Error))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected != 0 {
		err = ErrorExistFailed
		return
	}

	return
}

func CategoryCreate(name string) (err error) {
	category := model.Category{
		Name:   name,
		Speak:  0,
		Follow: 0,
	}

	if err = CategoryReapetCheck(category.Name); err != nil {
		zap.L().Error("insert user failed", zap.Error(err))
		return
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
		Type:       1,
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

func CategoryQuery(left int, right int) (category []model.Category, totNum int64, curNum int64, err error) {
	db := DatabaseGet()

	count := db.Limit(right - left).Offset(left).Find(&category)

	if count.Error != nil {
		err = ErrorQueryFailed
	}
	curNum = count.RowsAffected
	db.Model(&model.Category{}).Count(&totNum)

	return
}

func CategoryQueryById(categoryId uint) (category model.Category, err error) {

	db := DatabaseGet()

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
	return
}

func CategoryWikiSet(categoryId uint, postId uint) (err error) {
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

	category.WikiId = postId

	if err = db.Save(&category).Error; err != nil {
		db.Rollback()
		zap.L().Error("update post speak count failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

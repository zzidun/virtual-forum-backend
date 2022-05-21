package dao

import (
	"zzidun.tech/vforum0/model"
)

func CategoryerSet(categoryId uint, userId uint, categoryerType uint) (err error) {
	categoryer := model.Categoryer{
		CategoryId: categoryId,
		UserId:     userId,
		Type:       categoryerType,
	}

	if categoryerType == 1 {
		categoryerOld, err := CategoryerQueryByCategoryId(categoryId)
		if err != nil {
			return ErrorQueryFailed
		}
		if categoryerOld.ID != 0 {
			err = CategoryerCancel(categoryerOld.ID)
			if err != nil {
				return ErrorDeleteFailed
			}
		}
	}

	db := DatabaseGet()
	if err = db.Create(&categoryer).Error; err != nil {
		db.Rollback()
		err = ErrorInsertFailed
		return
	}

	return
}

func CategoryerCancel(categoryerId uint) (err error) {

	db := DatabaseGet()

	categoryer, err := CategoryerQueryById(categoryerId)

	if err = db.Delete(&categoryer).Error; err != nil {
		err = ErrorDeleteFailed
		return
	}

	return
}

func CategoryerQueryById(categoryerId uint) (categoryer *model.Categoryer, err error) {

	db := DatabaseGet()

	count := db.Where("id = ?", categoryerId).Find(&categoryer)

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

func CategoryerQueryByCategoryId(categoryId uint) (categoryer *model.Categoryer, err error) {

	db := DatabaseGet()

	count := db.Where("category_id = ? AND type = ?", categoryId, 1).Find(&categoryer)

	if count.Error != nil {
		err = ErrorQueryFailed
		return
	}

	return
}

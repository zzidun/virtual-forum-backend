package dao

import (
	"zzidun.tech/vforum0/model"
)

func CategoryReapetCheck(name string) (err error) {
	db := DatabaseGet()
	count := db.Where("name = ?", name).Find(&model.Category{})

	if count.Error != nil {
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
		return
	}

	db := DatabaseGet()

	if err = db.Create(&category).Error; err != nil {
		db.Rollback()
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

// 删除版块
func CategoryDelete(categoryId uint) (err error) {

	db := DatabaseGet()

	category, err := CategoryQueryById(categoryId)
	if err != nil {
		return
	}

	if err = db.Delete(&category).Error; err != nil {
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

func CategoryQueryById(categoryId uint) (category *model.Category, err error) {

	db := DatabaseGet()

	count := db.Where("id = ?", categoryId).Find(&category)

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

func CategoryWikiSet(categoryId uint, postId uint) (err error) {
	db := DatabaseGet()

	category, err := CategoryQueryById(categoryId)
	if err != nil {
		return
	}

	category.WikiId = postId

	if err = db.Save(&category).Error; err != nil {
		db.Rollback()
		err = ErrorInsertFailed
		return
	}

	return
}

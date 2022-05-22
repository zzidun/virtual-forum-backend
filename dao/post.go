package dao

import (
	"go.uber.org/zap"
	"zzidun.tech/vforum0/model"
)

func PostCreate(categoryId uint, title string, userId uint) (err error) {
	post := model.Post{
		CategoryId: categoryId,
		Title:      title,
		Speak:      0,
		UserId:     userId,
	}

	db := DatabaseGet()
	if err = db.Create(&post).Error; err != nil {
		db.Rollback()
		zap.L().Error("insert user failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

func PostDelete(postId uint) (err error) {

	db := DatabaseGet()

	post, err := PostQueryById(postId)
	if err != nil {
		return
	}

	if err = db.Delete(&post).Error; err != nil {
		zap.L().Error("delelte post failed", zap.Error(err))
		err = ErrorDeleteFailed
		return
	}

	return
}

func PostQueryByCategoryId(categoryId uint, left int, right int) (postList []model.Post, totNum int64, curNum int64, err error) {
	db := DatabaseGet()

	count := db.Where("category_id = ?", categoryId).Order("id desc").Limit(right - left).Offset(left).Find(&postList)

	if count.Error != nil {
		err = ErrorQueryFailed
	}
	curNum = count.RowsAffected
	db.Model(&model.Post{}).Where("category_id = ?", categoryId).Count(&totNum)

	return
}

func PostQueryByCategoryIdReplyTime(categoryId uint, left int, right int) (postList []model.Post, totNum int64, curNum int64, err error) {
	db := DatabaseGet()

	count := db.Where("category_id = ?", categoryId).Order("updated_at desc").Limit(right - left).Offset(left).Find(&postList)

	if count.Error != nil {
		err = ErrorQueryFailed
	}
	curNum = count.RowsAffected
	db.Model(&model.Post{}).Where("category_id = ?", categoryId).Count(&totNum)

	return
}

func PostQueryById(postId uint) (post *model.Post, err error) {
	db := DatabaseGet()

	count := db.Where("id = ?", postId).Find(&post)

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

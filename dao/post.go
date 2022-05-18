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

	var post model.Post

	count := db.Where("id = ?", postId).Find(&post)

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

	if err = db.Delete(&post).Error; err != nil {
		zap.L().Error("delelte post failed", zap.Error(err))
		err = ErrorDeleteFailed
		return
	}

	return
}

func PostQuery(postIdLeft uint, postIdRight uint) (err error) {
	return
}

func PostQueryById(postId uint) (err error) {
	return
}

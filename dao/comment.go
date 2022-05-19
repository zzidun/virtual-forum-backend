package dao

import (
	"go.uber.org/zap"
	"zzidun.tech/vforum0/model"
)

func CommentCreate(postId uint, userId uint, replyId uint, content string) (err error) {

	comment := model.Comment{
		PostId:  postId,
		UserId:  userId,
		ReplyId: replyId,
		Content: content,
	}

	db := DatabaseGet()
	if err = db.Create(&comment).Error; err != nil {
		db.Rollback()
		zap.L().Error("insert comment failed", zap.Error(err))
		err = ErrorInsertFailed
		return
	}

	return
}

func CommentDelete(commentId uint) (err error) {

	db := DatabaseGet()

	var comment model.Comment

	count := db.Where("id = ?", commentId).Find(&comment)

	if count.Error != nil {
		zap.L().Error("query comment failed", zap.Error(err))
		err = ErrorQueryFailed
		return
	}
	if count.RowsAffected == 0 {
		zap.L().Error("query comment failed", zap.Error(err))
		err = ErrorNotExistFailed
		return
	}

	if err = db.Delete(&comment).Error; err != nil {
		zap.L().Error("delelte comment failed", zap.Error(err))
		err = ErrorDeleteFailed
		return
	}

	return
}

func CommentQuery(commentIdLeft uint, commentIdRight uint) (err error) {
	return
}

func CommentQueryById(commentId uint) (err error) {
	return
}

func CommentQueryByPostId(portId uint) (commentList []model.Comment, err error) {
	db := DatabaseGet()

	if db.Where("post_id = ?", portId).Find(&commentList).Error != nil {
		err = ErrorQueryFailed
	}
	return
}

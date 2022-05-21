package dao

import (
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
		err = ErrorInsertFailed
		return
	}

	return
}

func CommentDelete(commentId uint) (err error) {

	db := DatabaseGet()

	comment, err := CommentQueryById(commentId)
	if err != nil {
		return
	}

	if err = db.Delete(&comment).Error; err != nil {
		db.Rollback()
		err = ErrorDeleteFailed
		return
	}

	return
}

func CommentQuery(commentIdLeft uint, commentIdRight uint) (err error) {
	return
}

func CommentQueryById(commentId uint) (comment *model.Comment, err error) {

	db := DatabaseGet()
	count := db.Where("id = ?", commentId).Find(&comment)

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

func CommentQueryByPostId(postId uint, left int, right int) (commentList []model.Comment, totNum int64, curNum int64, err error) {
	db := DatabaseGet()

	count := db.Where("post_id = ?", postId).Limit(right - left).Offset(left).Find(&commentList)

	if count.Error != nil {
		err = ErrorQueryFailed
	}
	curNum = count.RowsAffected
	db.Model(&model.Comment{}).Where("post_id = ?", postId).Count(&totNum)

	return
}

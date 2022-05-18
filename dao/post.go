package dao

import (
	"zzidun.tech/vforum0/model"
)

func PostCreate(categoryId uint, title string, userId uint) (err error) {
	return
}

func PostDelete(postId uint) (err error) {
	db := DatabaseGet()
	db.Where("id = ?", postId).Delete(&model.Post{})
	return
}

func PostQuery(postIdLeft uint, postIdRight uint) (err error) {
	return
}

func PostQueryById(postId uint) (err error) {
	return
}

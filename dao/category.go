package dao

import (
	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/model"
)

func CategoryCreate(ctx *gin.Context) (err error) {
	category := model.Category{
		Name:   "name",
		Speak:  0,
		Follow: 0,
	}

	db := DatabaseGet()
	db.Create(category)

	return
}

func CategoryerSet(ctx *gin.Context) (err error) {
	categoryer := model.Categoryer{
		CategoryId: 0,
		UserId:     0,
		AdminType:  true,
	}

	db := DatabaseGet()
	db.Create(categoryer)

	return
}

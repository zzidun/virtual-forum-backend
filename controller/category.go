package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"zzidun.tech/vforum0/model"
	"zzidun.tech/vforum0/response"
)

type ICategory_Controller interface {
	Rest_Contoller
}

type Category_Controller struct {
	DB *gorm.DB
}

func Category_Controller_New() ICategory_Controller {
	db := model.DB_Get()

	return Category_Controller{DB: db}
}

func (c Category_Controller) Insert(ctx *gin.Context) {
	var request_category model.Category

	ctx.Bind(&request_category)

	if request_category.Name == "" {
		response.Response_Fail_Make(ctx, nil, "数据验证错误")
		return
	}

	var category = model.Category{
		Name: request_category.Name,
	}
	c.DB.Create(&category)
	response.Response_Success_Make(ctx, gin.H{
		"category": category}, "")
	return
}

func (c Category_Controller) Remove(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	c.DB.Delete(&model.Category{}, id)

	response.Response_Success_Make(ctx, nil, "删除成功")
	return
}

func (c Category_Controller) Update(ctx *gin.Context) {
	name := ctx.PostForm("name")
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if name == "" {
		response.Response_Fail_Make(ctx, nil, "数据验证错误")
		return
	}
	var category model.Category
	c.DB.Where("ID = ?", id).First(&category)
	if category.ID == 0 {
		response.Response_Fail_Make(ctx, nil, "板块不存在")
		return
	}

	c.DB.Model(&category).Update("name", name)
	response.Response_Success_Make(ctx, gin.H{
		"category": category}, "")
	return
}

func (c Category_Controller) Query(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Params.ByName("id"))

	var category model.Category
	c.DB.Where("ID = ?", id).First(&category)
	if category.ID == 0 {
		response.Response_Fail_Make(ctx, nil, "板块不存在")
		return
	}

	response.Response_Success_Make(ctx, gin.H{
		"category": category}, "")
	return
}

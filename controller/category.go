package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/logic"
	"zzidun.tech/vforum0/model"
	"zzidun.tech/vforum0/response"
)

func CategoryCreate(ctx *gin.Context) {

	adminId, exist := ctx.Get("userId")
	if !exist {
		return
	}

	valid, err := AdminCheckPerm(adminId.(uint), CodeCategoryPerm)
	if err != nil || !valid {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	var ccform *model.CategoryCreateForm
	if err := ctx.ShouldBindJSON(&ccform); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SiginUp with invalid param", zap.Error(err))
		// 判断err是不是 validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			response.ResponseError(ctx, response.CodeInvalidParams) // 请求参数错误
			return
		}

		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, errs)
		return
	}

	if err := dao.CategoryCreate(ccform.Name); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	return
}

func CategoryDelete(ctx *gin.Context) {
	return
}

func CategoryQuery(ctx *gin.Context) {

	var clRequired *model.CategoryListRequired
	if err := ctx.ShouldBindJSON(&clRequired); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SiginUp with invalid param", zap.Error(err))
		// 判断err是不是 validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			response.ResponseError(ctx, response.CodeInvalidParams) // 请求参数错误
			return
		}

		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, errs)
		return
	}

	left, err := strconv.ParseInt(clRequired.Left, 10, 32)
	if err != nil {
		left = 1
	}
	right, err := strconv.ParseInt(clRequired.Right, 10, 32)
	if err != nil {
		right = 16
	}

	categoryList, err := logic.CategoryList(int(left), int(right))
	if err != nil {
		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, categoryList)
}

func CategoryQueryById(ctx *gin.Context) {
	return
}

func CategoryerSet(ctx *gin.Context) {
	return
}

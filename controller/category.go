package controller

import (
	"fmt"
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

	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeCategoryPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	fmt.Println(authId.(uint))

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

	response.Response(ctx, response.CodeSuccess, nil)

	return
}

func CategoryDelete(ctx *gin.Context) {

	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeCategoryPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	categoryIdStr := ctx.Param("id")

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	if err := dao.CategoryDelete(uint(categoryId)); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)

	return
}

func CategoryQuery(ctx *gin.Context) {

	leftStr := ctx.Query("left")
	rightStr := ctx.Query("right")

	left, err := strconv.ParseInt(leftStr, 10, 32)
	if err != nil {
		left = 0
	}

	right, err := strconv.ParseInt(rightStr, 10, 32)
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

	authId, exist := ctx.Get("authId")

	categoryIdStr := ctx.Param("id")

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	category, err := dao.CategoryQueryById(uint(categoryId))

	if err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	if exist {
		userFollow, err := dao.UserFollowQuery(authId.(uint), uint(categoryId))
		if err != nil {
			return
		}

		response.Response(ctx, response.CodeSuccess, gin.H{
			"followed": fmt.Sprintf("%d", userFollow.ID),
			"name":     category.Name,
			"speak":    fmt.Sprintf("%d", category.Speak),
			"follow":   fmt.Sprintf("%d", category.Follow),
			"wiki":     fmt.Sprintf("%d", category.WikiId),
		})
		return
	}

	response.Response(ctx, response.CodeSuccess, gin.H{
		"followed": "0",
		"name":     category.Name,
		"speak":    fmt.Sprintf("%d", category.Speak),
		"follow":   fmt.Sprintf("%d", category.Follow),
		"wiki":     fmt.Sprintf("%d", category.WikiId),
	})

	return
}

func CategoryUpdate(ctx *gin.Context) {

	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeCategoryPerm)
	fmt.Println(valid)

	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	categoryIdStr := ctx.Param("id")

	var cForm *model.CategoryerForm
	if err := ctx.ShouldBindJSON(&cForm); err != nil {
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

	userId, err := strconv.ParseInt(cForm.UserId, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "用户id错误")
		return
	}

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	categoryerType, err := strconv.ParseInt(cForm.Type, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版主类型错误")
		return
	}

	if err := dao.CategoryerSet(uint(categoryId), uint(userId), uint(categoryerType)); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)
	return
}

func CategoryWiki(ctx *gin.Context) {

	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	var cwForm *model.CategoryWikiForm
	if err := ctx.ShouldBindJSON(&cwForm); err != nil {
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

	categoryId, err := strconv.ParseInt(cwForm.CategoryId, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	postId, err := strconv.ParseInt(cwForm.PostId, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	valid, err := logic.CategoryerCheck(authId.(uint), uint(categoryId))
	if err != nil || !valid {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	if err := dao.CategoryWikiSet(uint(categoryId), uint(postId)); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)
	return
}

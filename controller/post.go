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

// 发布文章
func PostPost(ctx *gin.Context) {

	userId, exist := ctx.Get("userId")
	if !exist {
		return
	}

	var ppform *model.PostPostForm
	if err := ctx.ShouldBindJSON(&ppform); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SiginUp with invalid param", zap.Error(err))

		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, "参数错误")
		return
	}

	categoryId, err := strconv.ParseUint(ppform.CategoryId, 10, 32)
	if err != nil {
		zap.L().Error("SiginUp with invalid param", zap.Error(err))
		return
	}

	if err := dao.PostCreate(uint(categoryId), ppform.Title, userId.(uint)); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)
	return
}

func PostDelete(ctx *gin.Context) {

	return
}

func PostQuery(ctx *gin.Context) {

	var plRequired *model.PostListRequired
	if err := ctx.ShouldBindJSON(&plRequired); err != nil {
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

	left, err := strconv.ParseInt(plRequired.Left, 10, 32)
	if err != nil {
		left = 0
	}
	right, err := strconv.ParseInt(plRequired.Right, 10, 32)
	if err != nil {
		right = 15
	}

	categoryId, err := strconv.ParseInt(plRequired.CategoryId, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块不存在")
		return
	}

	postList, err := logic.PostList(uint(categoryId), int(left), int(right))
	if err != nil {
		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, postList)
}

func PostQueryReplyTime(ctx *gin.Context) {

	categoryId := 1
	postList, err := logic.PostList(uint(categoryId), 0, 15)
	if err != nil {
		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, postList)
}

func PostQueryById(ctx *gin.Context) {
	return
}

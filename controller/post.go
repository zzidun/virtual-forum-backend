package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/model"
	"zzidun.tech/vforum0/response"
)

func PostPost(ctx *gin.Context) {

	user_id, exist := ctx.Get("userId")
	if !exist {
		return
	}

	var ppform *model.PostPostForm
	if err := ctx.ShouldBindJSON(&ppform); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SiginUp with invalid param", zap.Error(err))
		// 判断err是不是 validator.ValidationErrors类型的errors
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			response.ResponseError(ctx, response.CodeInvalidParams) // 请求参数错误
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, errs)
		return
	}

	if fmt.Sprintf("%d", user_id.(uint)) != ppform.UserId {
		response.ResponseError(ctx, 100)
		return
	}

	if err := dao.PostCreate(ppform); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, 100)
		return
	}

	return
}

func PostDelete(ctx *gin.Context) {
	return
}

func PostQuery(ctx *gin.Context) {
	return
}

func PostQueryById(ctx *gin.Context) {
	return
}

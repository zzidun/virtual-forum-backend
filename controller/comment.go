package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/model"
	"zzidun.tech/vforum0/response"
)

func CommentPost(ctx *gin.Context) {

	user_id, exist := ctx.Get("userId")
	if !exist {
		return
	}

	var cpform *model.CommentPostForm
	if err := ctx.ShouldBindJSON(&cpform); err != nil {
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

	if fmt.Sprintf("%d", user_id.(uint)) != cpform.UserId {
		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	postId, err := strconv.ParseUint(cpform.PostId, 10, 32)
	if err != nil {
		zap.L().Error("SiginUp with invalid param", zap.Error(err))
		return
	}

	userId, err := strconv.ParseUint(cpform.UserId, 10, 32)
	if err != nil {
		zap.L().Error("SiginUp with invalid param", zap.Error(err))
		return
	}

	replyId, err := strconv.ParseUint(cpform.ReplyId, 10, 32)
	if err != nil {
		zap.L().Error("SiginUp with invalid param", zap.Error(err))
		return
	}

	if err := dao.CommentCreate(uint(postId), uint(userId), uint(replyId), cpform.Content); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	return
}

func CommentDelete(ctx *gin.Context) {
	return
}

func CommentQuery(ctx *gin.Context) {
	return
}

func CommentQueryById(ctx *gin.Context) {
	return
}

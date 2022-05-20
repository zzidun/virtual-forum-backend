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

func CommentPost(ctx *gin.Context) {

	userId, exist := ctx.Get("userId")
	if !exist {
		return
	}

	var cpform *model.CommentPostForm
	if err := ctx.ShouldBindJSON(&cpform); err != nil {
		// 请求参数有误，直接返回响应
		zap.L().Error("SiginUp with invalid param", zap.Error(err))

		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, "参数错误")
		return
	}

	postId, err := strconv.ParseUint(cpform.PostId, 10, 32)
	if err != nil {
		zap.L().Error("SiginUp with invalid param", zap.Error(err))
		response.Response(ctx, response.CodeUnknownError, "id解析错误")
		return
	}

	replyId, err := strconv.ParseUint(cpform.ReplyId, 10, 32)
	if err != nil {
		if cpform.ReplyId == "" {
			replyId = 0
		} else {
			zap.L().Error("SiginUp with invalid param", zap.Error(err))
			response.Response(ctx, response.CodeUnknownError, "回复id解析错误")
			return
		}
	}

	if err := dao.CommentCreate(uint(postId), userId.(uint), uint(replyId), cpform.Content); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)

	return
}

func CommentDelete(ctx *gin.Context) {
	return
}

func CommentQuery(ctx *gin.Context) {

	var clRequired *model.CommentListRequired
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

	postId, err := strconv.ParseInt(clRequired.PostId, 10, 32)
	if err != nil {
		right = 16
	}

	commentList, err := logic.CommentList(uint(postId), int(left), int(right))
	if err != nil {
		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, commentList)
}

func CommentQueryById(ctx *gin.Context) {
	return
}

package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
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

	leftStr := ctx.Query("left")
	rightStr := ctx.Query("right")
	postIdStr := ctx.Query("post")

	left, err := strconv.ParseInt(leftStr, 10, 32)
	if err != nil {
		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}
	right, err := strconv.ParseInt(rightStr, 10, 32)
	if err != nil {
		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}

	postId, err := strconv.ParseInt(postIdStr, 10, 32)
	if err != nil {
		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}

	commentList, err := logic.CommentList(uint(postId), int(left), int(right))
	if err != nil {
		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}

	fmt.Println(left, right, postId)

	response.Response(ctx, response.CodeSuccess, commentList)
}

func CommentQueryById(ctx *gin.Context) {
	return
}

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

// 发布文章
func PostPost(ctx *gin.Context) {

	authId, exist := ctx.Get("authId")
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

	if err := dao.PostCreate(uint(categoryId), ppform.Title, authId.(uint)); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	if err = dao.CategoryCountSpeak(uint(categoryId)); err != nil {
		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	if err = dao.UserCountSpeak(authId.(uint)); err != nil {
		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	if err = dao.UserCategoryCount(authId.(uint), uint(categoryId)); err != nil {
		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)
	return
}

func PostDelete(ctx *gin.Context) {

	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	postIdStr := ctx.Param("id")

	postId, err := strconv.ParseInt(postIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	post, err := dao.PostQueryById(uint(postId))
	if err != nil {
		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	valid, err := logic.CategoryerCheck(authId.(uint), post.CategoryId)
	if err != nil {
		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	if post.UserId != authId.(uint) && !valid {
		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	if err := dao.PostDelete(uint(postId)); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)

	return
}

func PostQuery(ctx *gin.Context) {

	leftStr := ctx.Query("left")
	rightStr := ctx.Query("right")
	categoryIdStr := ctx.Query("category")

	left, err := strconv.ParseInt(leftStr, 10, 32)
	if err != nil {
		left = 0
	}
	right, err := strconv.ParseInt(rightStr, 10, 32)
	if err != nil {
		right = 15
	}

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 32)
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

	leftStr := ctx.Query("left")
	rightStr := ctx.Query("right")
	categoryIdStr := ctx.Query("category")

	left, err := strconv.ParseInt(leftStr, 10, 32)
	if err != nil {
		left = 0
	}
	right, err := strconv.ParseInt(rightStr, 10, 32)
	if err != nil {
		right = 15
	}

	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块不存在")
		return
	}

	postList, err := logic.PostListReplyTime(uint(categoryId), int(left), int(right))
	if err != nil {
		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, postList)
}

func PostQueryById(ctx *gin.Context) {

	postIdStr := ctx.Param("id")

	postId, err := strconv.ParseInt(postIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	post, err := dao.PostQueryById(uint(postId))

	if err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.Response(ctx, response.CodeSuccess, gin.H{
		"id":       fmt.Sprintf("%d", post.ID),
		"title":    post.Title,
		"speak":    fmt.Sprintf("%d", post.Speak),
		"category": fmt.Sprintf("%d", post.CategoryId),
	})

	return
}

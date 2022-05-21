package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/response"
)

func CategoryFollowCreate(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	categoryIdStr := ctx.Query("category")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	if err := dao.UserFollowCreate(authId.(uint), uint(categoryId)); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.ResponseSuccess(ctx, nil)

	return
}

func CategoryFollowQuery(ctx *gin.Context) {
	return
}

func CategoryFollowById(ctx *gin.Context) {
	return
}

func CategoryFollowDelete(ctx *gin.Context) {

	followIdStr := ctx.Param("id")

	followId, err := strconv.ParseInt(followIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	if err := dao.UserFollowDelete(uint(followId)); err != nil {

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.ResponseSuccess(ctx, nil)

	return
}

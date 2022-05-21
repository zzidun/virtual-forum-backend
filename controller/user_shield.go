package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/model"
	"zzidun.tech/vforum0/response"
)

func UserShieldCreate(ctx *gin.Context) {
	userId, exist := ctx.Get("userId")
	if !exist {
		return
	}

	var usForm *model.UserShieldForm
	if err := ctx.ShouldBindJSON(&usForm); err != nil {
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

	shieldUserId, err := strconv.ParseInt(usForm.ShielUserId, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	if err := dao.UserShieldCreate(userId.(uint), uint(shieldUserId)); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.ResponseSuccess(ctx, nil)

	return
}

func UserShieldQuery(ctx *gin.Context) {
	return
}

func UserShieldQueryById(ctx *gin.Context) {
	return
}

func UserShieldDelete(ctx *gin.Context) {

	shieldIdStr := ctx.Param("id")

	shieldId, err := strconv.ParseInt(shieldIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	if err := dao.UserShieldDelete(uint(shieldId)); err != nil {

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.ResponseSuccess(ctx, nil)

	return
}

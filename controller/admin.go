package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/model"
	"zzidun.tech/vforum0/response"
	"zzidun.tech/vforum0/util"
)

func AdminLogin(ctx *gin.Context) {

	var alform *model.AdminLoginForm
	if err := ctx.ShouldBindJSON(&alform); err != nil {
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

	id, err := dao.AdminLogin(alform)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", alform.Name), zap.Error(err))

		response.ResponseError(ctx, 100)
		return
	}

	atoken, rtoken, err := util.TokenReleaseAccess(1, id, alform.Name)

	response.ResponseSuccess(ctx, gin.H{
		"user_id":   fmt.Sprintf("%d", id),
		"user_name": alform.Name,
		"atoken":    atoken,
		"rtoken":    rtoken,
	})
}

func AdminCreate(ctx *gin.Context) {
	return
}

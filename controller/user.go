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
	"zzidun.tech/vforum0/util"
)

// 处理用户注册
func UserRegister(ctx *gin.Context) {

	var urform *model.UserRegisterForm
	if err := ctx.ShouldBindJSON(&urform); err != nil {
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

	if err := dao.UserCreate(urform); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.ResponseSuccess(ctx, nil)
}

// 处理用户登陆
func UserLogin(ctx *gin.Context) {

	var ulform *model.UserLoginForm
	if err := ctx.ShouldBindJSON(&ulform); err != nil {
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

	userId, err := dao.UserLogin(ulform)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", ulform.Name), zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	token, err := util.TokenRelease(0, userId, ulform.Name)

	dao.UserUpdateLoginIpv4(userId, ctx.ClientIP())

	response.ResponseSuccess(ctx, gin.H{
		"user_id":   fmt.Sprintf("%d", userId),
		"user_name": ulform.Name,
		"token":     token,
	})
}

func UserQuery(ctx *gin.Context) {

	userId, exist := ctx.Get("userId")

	userIdStr := ctx.Param("id")

	userId2, err := strconv.ParseInt(userIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	user, err := dao.UserQueryById(uint(userId2))

	if err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	if exist {
		UserShield, err := dao.UserShieldQuery(userId.(uint), uint(userId2))
		if err != nil {
			return
		}

		response.Response(ctx, response.CodeSuccess, gin.H{
			"id":       user.ID,
			"shielded": fmt.Sprintf("%d", UserShield.ID),
			"name":     user.Name,
			"speak":    fmt.Sprintf("%d", user.Speak),
			"signal":   user.Signal,
			"lastip":   user.LastLoginIpv4,
			"lasttime": user.UpdatedAt,
		})

		return
	}

	response.Response(ctx, response.CodeSuccess, gin.H{
		"id":       user.ID,
		"shielded": "0",
		"name":     user.Name,
		"speak":    fmt.Sprintf("%d", user.Speak),
		"signal":   user.Signal,
		"lastip":   user.LastLoginIpv4,
		"lasttime": user.UpdatedAt,
	})

	return
}

func UserUpdate(ctx *gin.Context) {

	userId, exist := ctx.Get("userId")
	if !exist {
		return
	}

	var uuForm *model.UserUpdateForm
	if err := ctx.ShouldBindJSON(&uuForm); err != nil {
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

	if err := dao.UserUpdate(userId.(uint), uuForm.Email, uuForm.Password, uuForm.Signal); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.ResponseSuccess(ctx, nil)

	return
}

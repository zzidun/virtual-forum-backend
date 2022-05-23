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

	var urForm *model.UserRegisterForm
	if err := ctx.ShouldBindJSON(&urForm); err != nil {
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

	_, err := dao.UserCreate(urForm.Name, urForm.Email, urForm.Password)
	if err != nil {
		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.ResponseSuccess(ctx, nil)
}

// 处理用户登陆
func UserLogin(ctx *gin.Context) {

	var ulForm *model.UserLoginForm
	if err := ctx.ShouldBindJSON(&ulForm); err != nil {
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

	user, err := dao.UserLogin(ulForm.Email, ulForm.Password)
	if err != nil {
		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	token, err := util.TokenRelease(user.ID, user.Name)

	dao.UserUpdateLoginIpv4(user.ID, ctx.ClientIP())

	response.ResponseSuccess(ctx, gin.H{
		"user_id":   fmt.Sprintf("%d", user.ID),
		"user_name": user.Name,
		"token":     token,
	})
}

func UserQuery(ctx *gin.Context) {

	authId, exist := ctx.Get("authId")

	userIdStr := ctx.Param("id")

	userId, err := strconv.ParseInt(userIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	user, err := dao.UserQueryById(uint(userId))

	if err != nil {
		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	if exist {
		UserShield, err := dao.UserShieldQuery(authId.(uint), uint(userId))
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

	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	userIdStr := ctx.Param("id")

	userId, err := strconv.ParseInt(userIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
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

	if authId.(uint) != uint(userId) {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	_, err = dao.UserUpdate(authId.(uint), uuForm.Email, uuForm.Password, uuForm.Signal,
		uuForm.PasswordOld)

	if err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	response.ResponseSuccess(ctx, nil)

	return
}

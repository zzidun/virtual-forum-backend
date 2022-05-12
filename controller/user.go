package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/dto"
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

	if err := dao.UserRegister(urform); err != nil {
		zap.L().Error("logic.signup failed", zap.Error(err))

		response.ResponseError(ctx, 100)
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

	id, err := dao.UserLogin(ulform)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", ulform.Name), zap.Error(err))

		response.ResponseError(ctx, 100)
		return
	}

	atoken, rtoken, err := util.TokenReleaseAccess(0, id, ulform.Name)

	response.ResponseSuccess(ctx, gin.H{
		"user_id":   fmt.Sprintf("%d", id),
		"user_name": ulform.Name,
		"atoken":    atoken,
		"rtoken":    rtoken,
	})
}

func User_Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")

	response.Response_Success_Make(ctx, gin.H{"user": dto.UserDto_Make(user.(model.User))}, "")
	return
}

func Email_Exist(db *gorm.DB, email string) bool {
	var user model.User
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

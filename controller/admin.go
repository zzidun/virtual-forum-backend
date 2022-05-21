package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/logic"
	"zzidun.tech/vforum0/model"
	"zzidun.tech/vforum0/response"
)

func AdminCreate(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeAdminPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	var aForm *model.AdminForm
	if err := ctx.ShouldBindJSON(&aForm); err != nil {
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

	userId, err := strconv.ParseInt(aForm.UserId, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}
	adminPerm, err := strconv.ParseInt(aForm.AdminPerm, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}
	banPerm, err := strconv.ParseInt(aForm.BanPerm, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}
	categoryPerm, err := strconv.ParseInt(aForm.CategoryPerm, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	admin, err := dao.AdminCreate(uint(userId))
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	err = dao.AdminUpdate(admin.ID, uint(adminPerm), uint(banPerm), uint(categoryPerm))
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)

	return
}

func AdminDelete(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeAdminPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	userIdStr := ctx.Param("id")

	userId, err := strconv.ParseInt(userIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	if err = dao.AdminDelete(uint(userId)); err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)

	return
}

func AdminUpdate(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeAdminPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	userIdStr := ctx.Param("id")

	userId, err := strconv.ParseInt(userIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	var aForm *model.AdminForm
	if err := ctx.ShouldBindJSON(&aForm); err != nil {
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

	adminPerm, err := strconv.ParseInt(aForm.AdminPerm, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}
	banPerm, err := strconv.ParseInt(aForm.BanPerm, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}
	categoryPerm, err := strconv.ParseInt(aForm.CategoryPerm, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	user, err := dao.UserQueryById(uint(userId))

	err = dao.AdminUpdate(user.AdminId, uint(adminPerm), uint(banPerm), uint(categoryPerm))
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)

	return
}

func AdminQuery(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeAdminPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	leftStr := ctx.Query("left")
	rightStr := ctx.Query("right")

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

	adminList, err := logic.AdminList(int(left), int(right))

	if err != nil {
		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, adminList)

	return
}

func AdminQueryById(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeAdminPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	userIdStr := ctx.Param("id")

	userId, err := strconv.ParseInt(userIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	user, err := dao.UserQueryById(uint(userId))
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	admin, err := dao.AdminQueryById(user.AdminId)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	response.Response(ctx, response.CodeSuccess, gin.H{
		"userid":       fmt.Sprintf("%d", user.ID),
		"userName":     user.Name,
		"adminperm":    fmt.Sprintf("%d", admin.AdminPerm),
		"banperm":      fmt.Sprintf("%d", admin.BanPerm),
		"categoryperm": fmt.Sprintf("%d", admin.CategoryPerm),
	})

	return
}

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

type AdminPermCode int64

const (
	CodeAdminPerm    AdminPermCode = 1000
	CodeBanPerm      AdminPermCode = 1001
	CodeCategoryPerm AdminPermCode = 1002
	CodePostPerm     AdminPermCode = 1003
)

func AdminCheckPerm(adminId uint, code AdminPermCode) (valid bool, err error) {
	admin, err := dao.AdminQueryById(adminId)
	if err != nil {
		return false, err
	}

	admingroup, err := dao.AdmingroupQueryById(admin.GroupId)
	if err != nil {
		return false, err
	}

	switch code {
	case CodeAdminPerm:
		valid = admingroup.AdminPerm
	case CodeBanPerm:
		valid = admingroup.BanPerm
	case CodeCategoryPerm:
		valid = admingroup.CategoryPerm
	case CodePostPerm:
		valid = admingroup.PostPerm
	default:
		return false, nil
	}

	return
}

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

		response.ResponseError(ctx, response.CodeUnknownError)
		return
	}

	token, err := util.TokenRelease(1, id, alform.Name)

	response.ResponseSuccess(ctx, gin.H{
		"user_id":   fmt.Sprintf("%d", id),
		"user_name": alform.Name,
		"token":     token,
	})
}

func AdminCreate(ctx *gin.Context) {
	return
}

func AdminDelete(ctx *gin.Context) {
	return
}

func AdminUpdate(ctx *gin.Context) {
	return
}

func AdminQuery(ctx *gin.Context) {
	return
}

func AdminQueryById(ctx *gin.Context) {
	return
}

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

func BanCreate(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeBanPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	var bForm *model.BannedIpv4Form
	if err := ctx.ShouldBindJSON(&bForm); err != nil {
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

	userId, err := strconv.ParseInt(bForm.UserId, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	_, err = dao.BannedIpv4Create(bForm.LeftIp, bForm.RightIp, uint(userId))
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)

	return
}

func BanDelete(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeBanPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	banIdStr := ctx.Param("id")

	banId, err := strconv.ParseInt(banIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	if err = dao.BannedIpv4Delete(uint(banId)); err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)

	return
}

func BanUpdate(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeBanPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	banIdStr := ctx.Param("id")

	banId, err := strconv.ParseInt(banIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	var bForm *model.BannedIpv4Form
	if err := ctx.ShouldBindJSON(&bForm); err != nil {
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

	userId, err := strconv.ParseInt(bForm.UserId, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	err = dao.BannedIpv4Update(uint(banId), bForm.LeftIp, bForm.RightIp, uint(userId))
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, nil)

	return
}

func BanQuery(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeBanPerm)
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

	banList, err := logic.BanList(int(left), int(right))

	if err != nil {
		response.Response(ctx, response.CodeUnknownError, nil)
		return
	}

	response.Response(ctx, response.CodeSuccess, banList)

	return
}

func BanQueryById(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeBanPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	banIdStr := ctx.Param("id")

	banId, err := strconv.ParseInt(banIdStr, 10, 32)
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	ban, err := dao.BannedIpv4QueryById(uint(banId))
	if err != nil {
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
		return
	}

	response.Response(ctx, response.CodeSuccess, gin.H{
		"banid":  fmt.Sprintf("%d", ban.ID),
		"left":   ban.LeftIp,
		"right":  ban.RightIp,
		"userid": fmt.Sprintf("%d", ban.UserId),
	})

	return
}

func FailQuery(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeBanPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	// leftStr := ctx.Query("left")
	// rightStr := ctx.Query("right")

	// left, err := strconv.ParseInt(leftStr, 10, 32)
	// if err != nil {
	// 	response.Response(ctx, response.CodeUnknownError, nil)
	// 	return
	// }
	// right, err := strconv.ParseInt(rightStr, 10, 32)
	// if err != nil {
	// 	response.Response(ctx, response.CodeUnknownError, nil)
	// 	return
	// }

	// failedList, err := logic.FailedList(int(left), int(right))

	// if err != nil {
	// 	response.Response(ctx, response.CodeUnknownError, nil)
	// 	return
	// }

	// response.Response(ctx, response.CodeSuccess, failedList)

	return

}

func FailQueryById(ctx *gin.Context) {
	authId, exist := ctx.Get("authId")
	if !exist {
		return
	}

	valid, err := logic.UserCheckPerm(authId.(uint), logic.CodeBanPerm)
	if err != nil || valid == 0 {
		response.ResponseErrorWithMsg(ctx, response.CodeUnknownError, nil)
		return
	}

	// failedIdStr := ctx.Param("id")

	// failedId, err := strconv.ParseInt(failedIdStr, 10, 32)
	// if err != nil {
	// 	response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
	// 	return
	// }

	// failed, err := dao.FailedQueryById(uint(failedId))
	// if err != nil {
	// 	response.ResponseErrorWithMsg(ctx, response.CodeInvalidParams, "版块id错误")
	// 	return
	// }

	// response.Response(ctx, response.CodeSuccess, gin.H{
	// 	"banid":  fmt.Sprintf("%d", ban.ID),
	// 	"left":   ban.LeftIp,
	// 	"right":  ban.RightIp,
	// 	"userid": fmt.Sprintf("%d", ban.UserId),
	// })

	return

}

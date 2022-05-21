package controller

import (
	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/logic"
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
}

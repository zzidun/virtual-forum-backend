package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response_Make(ctx *gin.Context, status int, code int, data gin.H, msg string) {
	ctx.JSON(status, gin.H{
		"code":    code,
		"data":    data,
		"message": msg,
	})
}

func Response_Success_Make(ctx *gin.Context, data gin.H, msg string) {
	Response_Make(ctx, http.StatusOK, http.StatusOK, data, msg)
}

func Response_Fail_Make(ctx *gin.Context, data gin.H, msg string) {
	Response_Make(ctx, http.StatusBadRequest, http.StatusBadRequest, data, msg)
}

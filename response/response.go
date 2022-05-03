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

type ResponseData struct {
	Code    ResponseCode `json:"code"`
	Message string       `json:"message"`
	Data    interface{}  `json:"data,omitempty"` // omitempty当data为空时,不展示这个字段
}

func ResponseError(ctx *gin.Context, c ResponseCode) {
	rd := &ResponseData{
		Code:    c,
		Message: c.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseErrorWithMsg(ctx *gin.Context, code ResponseCode, data interface{}) {
	rd := &ResponseData{
		Code:    code,
		Message: code.Msg(),
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, rd)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	rd := &ResponseData{
		Code:    CodeSuccess,
		Message: CodeSuccess.Msg(),
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rd)
}

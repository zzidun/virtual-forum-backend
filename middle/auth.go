package middle

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/response"
	"zzidun.tech/vforum0/util"
)

// 登陆验证中间件
func AuthMiddle() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.ResponseErrorWithMsg(ctx, response.CodeInvalidToken, "请求头缺少Auth Token")
			ctx.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.ResponseErrorWithMsg(ctx, response.CodeInvalidToken, "Token格式不对")
			ctx.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := util.TokenParse(parts[1])
		if err != nil {
			fmt.Println(err)
			response.ResponseError(ctx, response.CodeInvalidToken)
			ctx.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		ctx.Set("userType", mc.UserType)
		ctx.Set("userId", mc.UserId)
		ctx.Next() // 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}

// 管理员权限检测中间件
func AdminCheckMiddle() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		user_type, exist := ctx.Get("userType")
		if !exist || user_type.(uint) != 1 {
			response.ResponseErrorWithMsg(ctx, response.CodeInvalidToken, "请重新登陆")
			return
		}

		ctx.Next() // 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}

// 管理员权限检测中间件
func UserCheckMiddle() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		user_type, exist := ctx.Get("userType")
		if !exist || user_type != 0 {
			response.ResponseErrorWithMsg(ctx, response.CodeInvalidToken, "请重新登陆")
			return
		}

		ctx.Next() // 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}

package middle

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/logic"
)

//返回类型：HandlerFunc
func IpAuthorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取ip
		ip := ctx.ClientIP()
		//查询数据库是否存在该ip
		valid, err := logic.IpCheck(ip)
		if err != nil || !valid {
			// 验证不通过，不再调用后续的函数处理
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "访问未授权",
				"data":    nil,
			})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			return
		} else {
			ctx.Next()
		}
	}
}

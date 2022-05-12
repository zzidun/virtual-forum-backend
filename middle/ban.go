package middle

import "github.com/gin-gonic/gin"

// ip过滤中间件
func BanMiddle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		return
	}
}

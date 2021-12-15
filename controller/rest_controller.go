package controller

import "github.com/gin-gonic/gin"

type Rest_Contoller interface {
	Insert(ctx *gin.Context)
	Remove(ctx *gin.Context)
	Update(ctx *gin.Context)
	Query(ctx *gin.Context)
}

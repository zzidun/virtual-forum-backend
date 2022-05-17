package controller

import (
	"github.com/gin-gonic/gin"
	"zzidun.tech/vforum0/response"
)

func Slo(ctx *gin.Context) {
	response.ResponseSuccess(ctx, gin.H{"msg": "slo called"})
}

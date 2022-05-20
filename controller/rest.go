package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"zzidun.tech/vforum0/response"
)

func Slo(ctx *gin.Context) {
	response.ResponseSuccess(ctx, gin.H{"msg": "slo called"})
}

func About(ctx *gin.Context) {
	about := viper.Get("about")
	response.Response(ctx, response.CodeSuccess, gin.H{
		"about": about.(string),
	})
}

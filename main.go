package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"zzidun.tech/vforum0/dao"
	"zzidun.tech/vforum0/router"
	"zzidun.tech/vforum0/util"
)

func main() {
	// 不取消注释,将会是Debug模式
	// gin.SetMode(gin.ReleaseMode)
	util.ConfigInit()
	dao.Init()
	// if err := redis.Init(); err != nil {
	// 	return
	// }
	r := gin.Default()
	r = router.RouteInit(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	} else {
		panic(r.Run())
	}
}

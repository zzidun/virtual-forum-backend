package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"zzidun.tech/vforum0/model"
	"zzidun.tech/vforum0/router"
	"zzidun.tech/vforum0/util"
)

func main() {
	util.Config_Init()
	model.DB_Init()
	r := gin.Default()
	r = router.RouteInit(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	} else {
		panic(r.Run())
	}
}

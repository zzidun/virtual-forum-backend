package router

import (
	"github.com/gin-gonic/gin"

	"zzidun.tech/vforum0/controller"
	"zzidun.tech/vforum0/middle"
)

func AdminRoute(r *gin.Engine) *gin.Engine {
	return r
}

func UserRoute(r *gin.Engine) *gin.Engine {
	return r
}

func RouteInit(r *gin.Engine) *gin.Engine {
	r.Use(middle.Cors_Middle())

	r.POST("/user/register", controller.User_Register)
	r.POST("/user/login", controller.User_Login)
	r.GET("/user/info", middle.Auth_Middle(), controller.User_Info)

	category_router := r.Group("/category")
	category_controller := controller.Category_Controller_New()
	category_router.POST("", category_controller.Insert)
	category_router.PUT("/:id", category_controller.Update)
	category_router.GET("/:id", category_controller.Query)
	category_router.DELETE("/:id", category_controller.Remove)

	return r
}

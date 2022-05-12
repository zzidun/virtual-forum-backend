package router

import (
	"github.com/gin-gonic/gin"

	"zzidun.tech/vforum0/middle"
	"zzidun.tech/vforum0/response"
)

func slo(c *gin.Context) {
	response.ResponseSuccess(c, gin.H{"msg": "slo called"})
}

// 需要验证管理员身份的路由
func AdminRoute(r *gin.Engine) *gin.Engine {

	r.POST("/adminlogin", slo)

	admin_router := r.Group("/admins")
	admin_router.POST("", slo)
	admin_router.DELETE("/:id", slo)
	admin_router.PUT("/:id", slo)
	admin_router.GET("/:id", slo)
	admin_router.GET("", slo)

	admingroup_router := r.Group("/admingroups")
	admingroup_router.POST("", slo)
	admingroup_router.DELETE("/:id", slo)
	admingroup_router.PUT("/:id", slo)
	admingroup_router.GET("/:id", slo)
	admingroup_router.GET("", slo)

	ban_router := r.Group("/bans")
	ban_router.POST("", slo)
	ban_router.DELETE("/:id", slo)
	ban_router.PUT("/:id", slo)
	ban_router.GET("/:id", slo)
	ban_router.GET("", slo)

	category_router := r.Group("/categories")
	category_router.POST("", slo)
	category_router.DELETE("/:id", slo)
	category_router.PUT("/:id", slo)
	category_router.GET("/:id", slo)
	category_router.GET("", slo)

	return r
}

// 需要验证用户身份的路由
func UserRoute(r *gin.Engine) *gin.Engine {

	r.POST("/register", slo)
	r.POST("/login", slo)
	r.POST("/post", slo)

	return r
}

func RouteInit(r *gin.Engine) *gin.Engine {

	r.Use(middle.Cors_Middle())

	r = AdminRoute(r)
	r = UserRoute(r)

	// category_router := r.Group("/category")
	// category_controller := controller.Category_Controller_New()
	// category_router.POST("", category_controller.Insert)
	// category_router.PUT("/:id", category_controller.Update)
	// category_router.GET("/:id", category_controller.Query)
	// category_router.DELETE("/:id", category_controller.Remove)

	return r
}

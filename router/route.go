package router

import (
	"github.com/gin-gonic/gin"

	"zzidun.tech/vforum0/controller"
	"zzidun.tech/vforum0/middle"
	"zzidun.tech/vforum0/response"
)

func slo(ctx *gin.Context) {
	response.ResponseSuccess(ctx, gin.H{"msg": "slo called"})
}

// 需要验证管理员身份的路由
func AdminRoute(r *gin.Engine) *gin.Engine {

	// 管理员登陆
	r.POST("/adminlogin", controller.AdminLogin)

	admin_router := r.Group("/admins", middle.AuthMiddle())
	// 创建管理员
	admin_router.POST("", slo)
	// 删除管理员
	admin_router.DELETE("/:id", slo)
	// 修改管理员组别和密码
	admin_router.PUT("/:id", slo)
	// 获取当个管理员信息
	admin_router.GET("/:id", slo)
	// 获取管理员列表
	admin_router.GET("", slo)

	admingroup_router := r.Group("/admingroups", middle.AuthMiddle())
	admingroup_router.POST("", slo)
	admingroup_router.DELETE("/:id", slo)
	admingroup_router.PUT("/:id", slo)
	admingroup_router.GET("/:id", slo)
	admingroup_router.GET("", slo)

	ban_router := r.Group("/bans", middle.AuthMiddle())
	ban_router.POST("", slo)
	ban_router.DELETE("/:id", slo)
	ban_router.PUT("/:id", slo)
	ban_router.GET("/:id", slo)
	ban_router.GET("", slo)

	category_router := r.Group("/categories", middle.AuthMiddle())
	category_router.POST("", slo)
	category_router.DELETE("/:id", slo)
	category_router.PUT("/:id", slo)
	category_router.GET("/:id", slo)
	category_router.GET("", slo)

	return r
}

func ViewRouter(r *gin.Engine) *gin.Engine {
	// 获取版块列表
	r.GET("/categories")
	// 按id获取版块信息
	r.GET("/categories/:id")
	// 获取版块帖子列表
	r.GET("/posts")
	// 按id获取帖子信息
	r.GET("/posts/:id")
	// 获取评论列表
	r.GET("/comments")
	// 获取评论信息
	r.GET("/comments/:id")
	// 获取用户信息
	r.GET("/user:id")

	return r
}

// 需要验证用户身份的路由
func UserRoute(r *gin.Engine) *gin.Engine {

	// 用户注册
	r.POST("/register", controller.UserRegister)
	// 用户登陆
	r.POST("/login", controller.UserLogin)
	// 发布帖子
	r.POST("/posts")
	// 删除帖子
	r.DELETE("/posts/:id")
	// 发布评论
	r.POST("/comments")
	// 删除评论
	r.DELETE("/comments/:id")
	// 编辑用户信息
	r.PUT("/users/:id")
	// 屏蔽用户
	r.PUT("/shields")
	r.GET("/shields")
	r.GET("/shields/:id")
	r.DELETE("/shields/:id")
	// 收藏帖子
	r.PUT("/collects")
	r.GET("/collects")
	r.GET("/collects/:id")
	r.DELETE("/collects/:id")
	// 关注版块
	r.PUT("/follows")
	r.GET("/follows")
	r.GET("/follows/:id")
	r.DELETE("/follows/:id")

	return r
}

func RouteInit(r *gin.Engine) *gin.Engine {

	r.Use(middle.CorsMiddle())
	r.Use(middle.BanMiddle())

	r = AdminRoute(r)
	r = ViewRouter(r)
	r = UserRoute(r)

	return r
}

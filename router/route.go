package router

import (
	"github.com/gin-gonic/gin"

	"zzidun.tech/vforum0/controller"
	"zzidun.tech/vforum0/middle"
)

// 需要验证管理员身份的路由
func AdminRoute(r *gin.Engine) *gin.Engine {

	admin_router := r.Group("/admins", middle.AuthMiddle())
	// 创建管理员
	admin_router.POST("", controller.AdminCreate)
	// 删除管理员
	admin_router.DELETE("/:id", controller.AdminDelete)
	// 修改管理员权限
	admin_router.PUT("/:id", controller.AdminUpdate)
	// 获取单个管理员信息
	admin_router.GET("/:id", controller.AdminQueryById)
	// 获取管理员列表
	admin_router.GET("", controller.AdminQuery)

	ban_router := r.Group("/bans", middle.AuthMiddle())
	ban_router.POST("", controller.BanCreate)
	ban_router.DELETE("/:id", controller.BanDelete)
	ban_router.PUT("/:id", controller.BanUpdate)
	ban_router.GET("/:id", controller.BanQueryById)
	ban_router.GET("", controller.BanQuery)

	category_router := r.Group("/categories", middle.AuthMiddle())
	category_router.POST("", controller.CategoryCreate)
	category_router.DELETE("/:id", controller.CategoryDelete)
	category_router.PUT("/:id", controller.CategoryUpdate)

	r.GET("fails", controller.FailQuery)
	r.GET("fails/:id", controller.FailQueryById)

	return r
}

func ViewRouter(r *gin.Engine) *gin.Engine {
	// 获取论坛首页
	r.GET("/about", middle.IdMiddle(), controller.About)
	// 获取版块列表
	r.GET("/categories", middle.IdMiddle(), controller.CategoryQuery)
	// 按id获取版块信息
	r.GET("/categories/:id", middle.IdMiddle(), controller.CategoryQueryById)
	// 获取版块帖子列表
	r.GET("/posts", middle.IdMiddle(), controller.PostQuery)
	// 获取版块帖子列表
	r.GET("/posts2", middle.IdMiddle(), controller.PostQueryReplyTime)
	// 按id获取帖子信息
	r.GET("/posts/:id", middle.IdMiddle(), controller.PostQueryById)
	// 获取评论列表
	r.GET("/comments", middle.IdMiddle(), controller.CommentQuery)
	// 获取评论信息
	r.GET("/comments/:id", middle.IdMiddle(), controller.CommentQueryById)
	// 获取用户信息
	r.GET("/user/:id", middle.IdMiddle(), controller.UserQuery)

	return r
}

// 需要验证用户身份的路由
func UserRoute(r *gin.Engine) *gin.Engine {

	// 用户注册
	r.POST("/register", controller.UserRegister)
	// 用户登陆
	r.POST("/login", controller.UserLogin)
	// 发布帖子
	r.POST("/posts", middle.AuthMiddle(), controller.PostPost)
	// 删除帖子
	r.DELETE("/posts/:id", middle.AuthMiddle(), controller.PostDelete)
	// 发布评论
	r.POST("/comments", middle.AuthMiddle(), controller.CommentPost)
	// 删除评论
	r.DELETE("/comments/:id", middle.AuthMiddle(), controller.CommentDelete)
	// 编辑用户信息
	r.PUT("/users/:id", middle.AuthMiddle(), controller.UserUpdate)
	// 屏蔽用户
	r.POST("/shields", middle.AuthMiddle(), controller.UserShieldCreate)
	r.GET("/shields", middle.AuthMiddle(), controller.UserShieldQuery)
	r.GET("/shields/:id", middle.AuthMiddle(), controller.UserShieldQueryById)
	r.DELETE("/shields/:id", middle.AuthMiddle(), controller.UserShieldDelete)
	// 收藏帖子
	r.PUT("/collects", middle.AuthMiddle(), controller.PostCollectCreate)
	r.GET("/collects", middle.AuthMiddle(), controller.PostCollectQuery)
	r.GET("/collects/:id", middle.AuthMiddle(), controller.PostCollectQueryById)
	r.DELETE("/collects/:id", middle.AuthMiddle(), controller.PostCollectDelete)
	// 关注版块
	r.POST("/follows", middle.AuthMiddle(), controller.CategoryFollowCreate)
	r.GET("/follows", middle.AuthMiddle(), controller.CategoryFollowQuery)
	r.GET("/follows/:id", middle.AuthMiddle(), controller.CategoryFollowById)
	r.DELETE("/follows/:id", middle.AuthMiddle(), controller.CategoryFollowDelete)

	// 编辑版块信息
	r.POST("/wiki", middle.AuthMiddle(), controller.CategoryWiki)

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

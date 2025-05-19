package routes

import (
	"gateway/internal/handler"
	"gateway/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.ErrorMiddleware())
	// 公共路由组 - 不需要JWT验证
	public := ginRouter.Group("/api/v1")
	{
		public.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		public.POST("/user/register", handler.UserRegister)
		public.POST("/user/login", handler.UserLogin)
	}

	// 私有路由组 - 需要JWT验证
	private := public.Group("")
	private.Use(middleware.JWT()) // 应用JWT中间件
	{
		private.GET("/role/selRole", handler.SelRole)
		private.POST("/role/newAdminAPL", handler.NewAdminAPL)
		private.POST("/role/revAdminAPL", handler.RevAdminAPL)
		private.GET("/role/selAdminAPL", handler.SelAdminAPL)
	}

	// 权限路由组 - 需要JWT验证和权限验证
	rbac := private.Group("")
	rbac.Use(middleware.RBAC())
	{

	}

	return ginRouter
}

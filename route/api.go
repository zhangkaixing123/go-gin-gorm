package route

import (
	"github.com/gin-gonic/gin"
	"novel_spider/internal/handler"
	"novel_spider/middleware"
)

// RegisterApiRoutes 注册api路由
func RegisterApiRoutes(router *gin.Engine) {
	// 加载防跨域中间件
	router.Use(middleware.Cors())

	api := router.Group("/api")
	{
		api.GET("/test", handler.Test)
	}
}

package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"novel_spider/config"
	"novel_spider/middleware"
	router2 "novel_spider/route"
)

func New() {
	// 设置gin模式
	gin.SetMode(config.GetConf().Server.Mode)

	router := gin.Default()

	router.Use(middleware.Recover)

	// 读取自定义路由
	setRoute(router)

	if err := router.Run(fmt.Sprintf("%s:%s", config.GetConf().Server.Host, config.GetConf().Server.Port)); err != nil {
		panic(err)
	}
}

// 注册路由方法
func setRoute(r *gin.Engine) {
	router2.RegisterApiRoutes(r)
}

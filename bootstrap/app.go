package bootstrap

import (
	"novel_spider/config"
	"novel_spider/cron"
	"novel_spider/pkg/database"
	"novel_spider/pkg/gin"
	"novel_spider/pkg/logger"
	"novel_spider/pkg/redis"
)

func Start() {
	// 初始化log
	setUpLogger()

	// 初始化各种数据库
	redis.InitRedis()    // 初始化redis
	database.InitMysql() // 初始化MySQL

	// 启动定时任务
	_, f, err := cron.PowerServerProvider()
	if err != nil {
		// 如果报错，则关闭这个定时任务
		f()
	}

	// 启动gin
	gin.New()
}

// 初始化日志方法
func setUpLogger() {
	logger.InitLogger(
		config.GetConf().Logs.FileName,
		config.GetConf().Logs.MaxSize,
		config.GetConf().Logs.MaxBackup,
		config.GetConf().Logs.MaxAge,
		config.GetConf().Logs.Compress,
		config.GetConf().Logs.Type,
		config.GetConf().Logs.Level,
	)
}

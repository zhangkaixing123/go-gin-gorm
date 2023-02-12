package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"novel_spider/pkg/logger"
	"novel_spider/pkg/response"
	"runtime/debug"
)

// Recover 全局异常捕获
func Recover(cxt *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			// 将异常输出到日志
			var errorString = errorToString(r)
			logger.Logger.Error(errorString)
			debug.PrintStack()
			response.FailResponse(http.StatusInternalServerError, errorString).ToJson(cxt)
			cxt.Abort()
		}
	}()
	cxt.Next()
}

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}

package handler

import (
	"github.com/gin-gonic/gin"
	"novel_spider/internal/api"
	"novel_spider/internal/service"
	"novel_spider/pkg/response"
)

var Service service.Service

func Test(cxt *gin.Context) {
	var req api.TestReq

	// 将请求参数绑定到结构体里面
	if err := cxt.ShouldBind(&req); err != nil {
		response.ErrorResponse(400, "参数获取失败！").WriteTo(cxt)
		return
	}

	// 请求Service层的方法
	resp := Service.Test(cxt, &req)
	resp.WriteTo(cxt)
	return
}

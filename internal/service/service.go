package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"novel_spider/internal/api"
	"novel_spider/internal/dao"
	"novel_spider/pkg/response"
)

type Service struct {
}

var Dao = dao.Dao{}

func (s *Service) Test(cxt *gin.Context, req *api.TestReq) *response.JsonResponse {
	if err := req.Verify(); err != nil {
		return response.FailResponse(400, err.Error())
	}

	data := &api.TestResp{
		Name: req.Name,
		Page: req.Page,
	}
	return response.FailResponse(400, "密码错误", data)
}

func (s *Service) TimerTest() {
	fmt.Println("执行定时任务")
}

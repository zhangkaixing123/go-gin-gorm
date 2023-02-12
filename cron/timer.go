package cron

import (
	"github.com/robfig/cron"
	"novel_spider/internal/service"
)

// ┌─────────────second 范围 (0 - 60)
// │ ┌───────────── min (0 - 59)
// │ │ ┌────────────── hour (0 - 23)
// │ │ │ ┌─────────────── day of month (1 - 31)
// │ │ │ │ ┌──────────────── month (1 - 12)
// │ │ │ │ │ ┌───────────────── day of week (0 - 6) (0 to 6 are Sunday to
// │ │ │ │ │ │                  Saturday)
// │ │ │ │ │ │
// │ │ │ │ │ │
// * * * * * *
const (
	// 每天00:00
	CronSpecAt0000 = "0 0 0 * * ?"
	// 每天00:30
	CronSpecAt0030 = "0 30 0 * * ?"
	// 每天01:00
	CronSpecAt0100 = "0 0 1 * * *"
	// 每天02:00
	CronSpecAt0200 = "0 0 2 * * *"
	// 每天02:30
	CronSpecAt0230 = "0 30 2 * * *"
	// 每天02:50
	CronSpecAt0250 = "50 2 * * *"
	// 每天03:00
	CronSpecAt0300 = "0 3 * * *"
	// 每天06:00
	CronSpecAt0600 = "0 6 * * *"
	// 每天22:00
	CronSpecAt2200 = "0 22 * * *"
	CronSpecAt2300 = "0 23 * * *"
	// 每周一02:00
	CronSpecAt0200Mon = "0 2 * * 1"
	// 每周一05:00
	CronSpecAt0500Mon = "0 5 * * 1"
	// 每周一06:00
	CronSpecAt0600Mon = "0 6 * * 1"
	// 每隔60分钟
	CronSpecEvery60thMin = "0 0/60 * * * ?"
	// 每隔5分钟
	CronSpecEvery5thMin = "0 0/5 * * * ?"
	// 每隔30秒
	CronSpecEvery30 = "0/30 * * * * ?"
	// 每隔5秒
	CronSpecEvery5 = "0/5 * * * * ?"
	// 每月1号凌晨0点执行一次
	CronSpecEvery10 = "0/10 * * * * ?"
	// 测试使用
	CronA = "0/5 * * * * ?"
	// 测试使用
	CronB = "0/1 * * * * ?"
)

type Server struct {
	cron *cron.Cron
	svc  *service.Service
}

func NewServer(svc *service.Service) *Server {
	c := cron.New()
	return &Server{cron: c, svc: svc}
}

func (s *Server) Start() {
	s.cron.Start()
}

func (s *Server) Close() {
	s.cron.Stop()
}

func PowerServerProvider() (s *Server, cf func(), err error) {
	svc := &service.Service{}
	s = NewServer(svc)

	cf = s.Close

	if err = s.CronTasks(); err != nil {
		return
	}

	s.Start()

	return
}

func (s *Server) CronTasks() error {
	for _, j := range []struct {
		Name string
		Spec string
		Cmd  func()
	}{
		//{
		//	Name: "测试使用",
		//	Spec: func() string {
		//		return CronA
		//	}(),
		//	Cmd: s.svc.TimerTest,
		//},
	} {
		if err := s.cron.AddFunc(j.Spec, j.Cmd); err != nil {
			return err
		}
	}
	return nil
}

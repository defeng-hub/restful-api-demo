package impl

import (
	"github.com/infraboard/mcube/logger/zap"
	"restful-api-demo/apps/host"
	"restful-api-demo/conf"
)

// Config #####通过实现了下边两个方法就可以注册到ioc层了#####
func (s *MysqlServiceImpl) Config() {
	// 只需要保证config() 执行完成就能实现初始化
	s.db, _ = conf.C().MySQL.GetDB()
	s.l = zap.L().Named("Host")
}

func (s *MysqlServiceImpl) Name() string {
	return host.AppName
}

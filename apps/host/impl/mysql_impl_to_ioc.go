package impl

import (
	"github.com/infraboard/mcube/logger/zap"
	"restful-api-demo/apps"
	"restful-api-demo/apps/host"
	"restful-api-demo/conf"
)

// 这个对象要进入ioc模块,不过他当前并没有准备好, 需要对实例执行config方法
var impl = &MysqlServiceImpl{}

// Config #####通过实现了下边两个方法就可以注册到ioc层了#####
func (s *MysqlServiceImpl) Config() {
	// 只需要保证config() 执行完成就能实现初始化
	s.db, _ = conf.C().MySQL.GetDB()
	s.l = zap.L().Named("Host")
}

func (s *MysqlServiceImpl) Name() string {
	return host.AppName
}

func init() {
	// 老方法都是在start的时候,手动把服务注册到IOC层,  案例: apps.HostService = impl.NewMysqlServiceImpl()
	// 现在采用自动注册,类似于mysql引擎   import _ "xxx"
	// sql这个库就是案例
	apps.RegistryImpl(impl)
}

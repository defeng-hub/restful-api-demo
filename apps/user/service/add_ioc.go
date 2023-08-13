package service

import (
	"github.com/infraboard/mcube/logger/zap"
	"restful-api-demo/apps/user"
	"restful-api-demo/conf"
)

var ImplMap = map[string]string{
	"sys_user": "UserService",
	"sys_menu": "MenuService",
}

// Config 通过实现了下边两个方法就可以注册到ioc层了
// sys_user.go
func (s *UserService) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
	s.l = zap.L().Named("User Service")
}
func (s *UserService) Name() string {
	return user.AppName + ImplMap["sys_user"]
}

// Config 通过实现了下边两个方法就可以注册到ioc层了
// sys_menu.go
func (s *MenuService) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
}
func (s *MenuService) Name() string {
	return user.AppName + ImplMap["sys_menu"]
}

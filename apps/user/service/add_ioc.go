package system

import (
	"github.com/infraboard/mcube/logger/zap"
	"restful-api-demo/conf"
)

// Config 通过实现了下边两个方法就可以注册到ioc层了
// sys_user.go
func (s *UserService) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
	s.l = zap.L().Named("User Service")
}
func (s *UserService) Name() string {
	return "user.sys_user"
}

// Config 通过实现了下边两个方法就可以注册到ioc层了
// sys_menu.go
func (s *MenuService) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
}
func (s *MenuService) Name() string {
	return "user.sys_menu"
}

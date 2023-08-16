package service

import (
	"restful-api-demo/apps/user"
	"restful-api-demo/conf"
)

var ImplMap = map[string]string{
	"sys_user":   "UserService",
	"sys_menu":   "MenuService",
	"sys_casbin": "CasbinService",
}

// TODO:需要在这注册每一个app

// Config 通过实现了下边两个方法就可以注册到ioc层了
// sys_user.go
func (s *UserService) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
	s.l = conf.L().Named(s.Name())
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

// Name sys_casbin.go
func (s *CasbinService) Name() string {
	return user.AppName + ImplMap["sys_casbin"]
}
func (s *CasbinService) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
}

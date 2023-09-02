package add_ioc

import (
	"restful-api-demo/apps"
	svr "restful-api-demo/apps/user/service"
)

// 不能使用 var userService *svr.UserService，这种方式，会panic空指针异常，因为没有分配具体的内存，相当于只定义了一个类型
var userService = new(svr.UserService)
var menuService = new(svr.MenuService)
var casbinService = new(svr.CasbinService)
var baseMenuService = new(svr.BaseMenuService)
var authorityService = new(svr.AuthorityService)

func init() {
	apps.RegistryApp(userService)
	apps.RegistryApp(menuService)
	apps.RegistryApp(casbinService)
	apps.RegistryApp(baseMenuService)
	apps.RegistryApp(authorityService)
}

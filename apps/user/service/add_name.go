package service

var ImplMap = map[string]string{
	"sys_user":      "UserService",
	"sys_menu":      "MenuService",
	"base_sys_menu": "BaseMenuService",
	"sys_casbin":    "CasbinService",
	"sys_authority": "AuthorityService",
}

// TODO:需要在这注册每一个app

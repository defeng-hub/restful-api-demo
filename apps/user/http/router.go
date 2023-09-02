package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/user/http/middleware"
)

var Private = []gin.HandlerFunc{middleware.JWTAuth(), middleware.CasbinHandler(), middleware.OperationRecord()}
var Public = []gin.HandlerFunc{}

func Router(h *Handler, r gin.IRouter) {

	//用户模块
	RouterUserApi := r.Group("UserService").Use(Public...)
	RouterCasbinUserApi := r.Group("UserService").Use(Private...)
	{
		RouterUserApi.POST("Login", h.UserApi.Login)
		RouterUserApi.POST("Register", h.UserApi.Register)             // 用户注册账号
		RouterUserApi.POST("ChangePassword", h.UserApi.ChangePassword) // 用户修改密码

		RouterCasbinUserApi.PUT("SetSelfInfo", h.UserApi.SetSelfInfo)                // 设置自身信息
		RouterCasbinUserApi.GET("GetUserInfo", h.UserApi.GetUserInfo)                // 获取自身信息
		RouterCasbinUserApi.POST("GetUserList", h.UserApi.GetUserList)               // 分页获取用户列表
		RouterCasbinUserApi.POST("SetUserAuthorities", h.UserApi.SetUserAuthorities) // 设置用户权限组
		RouterCasbinUserApi.POST("ResetPassword", h.UserApi.ResetPassword)           // 设置用户权限组
		RouterCasbinUserApi.POST("SetUserAuthority", h.UserApi.SetUserAuthority)     // 设置用户权限
		RouterCasbinUserApi.DELETE("DeleteUser", h.UserApi.DeleteUser)               // 删除用户
		RouterCasbinUserApi.PUT("SetUserInfo", h.UserApi.SetUserInfo)                // 设置用户信息
	}

	//权限控制模块, 已完结
	RouterCasbinApi := r.Group("CasbinService").Use(Public...) //todo:待改回私有
	{
		RouterCasbinApi.POST("GetPolicyPathByAuthorityId", h.CasbinApi.GetPolicyPathByAuthorityId)
		RouterCasbinApi.POST("UpdateCasbin", h.CasbinApi.UpdateCasbin)
	}

	// Menu 菜单管理
	menuRouter := r.Group("MenuService").Use(Private...)
	menuRouterWithoutRecord := r.Group("MenuService").Use(Public...)
	{
		menuRouter.POST("addBaseMenu", h.MenuApi.AddBaseMenu)           // 新增菜单
		menuRouter.POST("addMenuAuthority", h.MenuApi.AddMenuAuthority) // 增加menu和角色关联关系
		menuRouter.POST("deleteBaseMenu", h.MenuApi.DeleteBaseMenu)     // 删除菜单
		menuRouter.POST("updateBaseMenu", h.MenuApi.UpdateBaseMenu)     // 更新菜单
	}
	{
		menuRouterWithoutRecord.POST("getMenu", h.MenuApi.GetMenu)                   // 获取菜单树
		menuRouterWithoutRecord.POST("getMenuList", h.MenuApi.GetMenuList)           // 分页获取基础menu列表
		menuRouterWithoutRecord.POST("getBaseMenuTree", h.MenuApi.GetBaseMenuTree)   // 获取用户动态路由
		menuRouterWithoutRecord.POST("getMenuAuthority", h.MenuApi.GetMenuAuthority) // 获取指定角色menu
		menuRouterWithoutRecord.POST("getBaseMenuById", h.MenuApi.GetBaseMenuById)   // 根据id获取菜单
	}
}

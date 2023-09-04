package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/user/http/middleware"
)

//middleware.OperationRecord()
var Private = []gin.HandlerFunc{middleware.JWTAuth(), middleware.CasbinHandler()}

//var Private = []gin.HandlerFunc{}
var Public = []gin.HandlerFunc{}

func Router(h *Handler, r gin.IRouter) {

	//用户模块
	RouterUserApi := r.Group("UserService").Use(Public...)
	RouterCasbinUserApi := r.Group("UserService").Use(Private...)
	{
		RouterUserApi.POST("Login", h.UserApi.Login)
		RouterUserApi.POST("Register", h.UserApi.Register)             // 用户注册账号
		RouterUserApi.POST("ChangePassword", h.UserApi.ChangePassword) // 用户修改密码
		RouterUserApi.POST("Logout", h.UserApi.JsonInBlacklist)        // 设置用户信息

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
	RouterCasbinApi := r.Group("CasbinService").Use(Private...) //todo:待改回私有
	{
		RouterCasbinApi.POST("GetPolicyPathByAuthorityId", h.CasbinApi.GetPolicyPathByAuthorityId)
		RouterCasbinApi.POST("UpdateCasbin", h.CasbinApi.UpdateCasbin)
	}

	// 权限管理
	authorityRouter := r.Group("AuthorityService").Use(Private...)
	{
		authorityRouter.POST("createAuthority", h.AuthorityApi.CreateAuthority)   // 创建角色
		authorityRouter.POST("deleteAuthority", h.AuthorityApi.DeleteAuthority)   // 删除角色
		authorityRouter.PUT("updateAuthority", h.AuthorityApi.UpdateAuthority)    // 更新角色
		authorityRouter.POST("copyAuthority", h.AuthorityApi.CopyAuthority)       // 拷贝角色
		authorityRouter.POST("setDataAuthority", h.AuthorityApi.SetDataAuthority) // 设置角色资源权限
	}
	{
		authorityRouter.POST("getAuthorityList", h.AuthorityApi.GetAuthorityList) // 获取角色列表
	}

	// Api管理
	apiRouter := r.Group("ApiService").Use(Private...)
	{
		apiRouter.POST("createApi", h.SystemApiApi.CreateApi)               // 创建Api
		apiRouter.POST("deleteApi", h.SystemApiApi.DeleteApi)               // 删除Api
		apiRouter.POST("getApiById", h.SystemApiApi.GetApiById)             // 获取单条Api消息
		apiRouter.POST("updateApi", h.SystemApiApi.UpdateApi)               // 更新api
		apiRouter.DELETE("deleteApisByIds", h.SystemApiApi.DeleteApisByIds) // 删除选中api
	}
	{
		apiRouter.POST("getAllApis", h.SystemApiApi.GetAllApis)          // 获取所有api
		apiRouter.POST("getApiList", h.SystemApiApi.GetApiList)          // 获取Api列表
		apiRouter.GET("getAllApiGroups", h.SystemApiApi.GetAllApiGroups) // 获取Api列表
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

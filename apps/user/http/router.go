package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/user/http/middleware"
	"restful-api-demo/apps/user/service"
)

func Router(h *Handler, r gin.IRouter) {

	//用户模块
	RouterUserApi := r.Group(service.ImplMap["sys_user"]).Use(middleware.OperationRecord())
	RouterCasbinUserApi := r.Group(service.ImplMap["sys_user"]).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		RouterUserApi.POST("Login", h.UserApi.Login)
		RouterUserApi.POST("Register", h.UserApi.Register)      // 用户注册账号
		RouterUserApi.PUT("SetSelfInfo", h.UserApi.SetSelfInfo) // 设置自身信息
		RouterUserApi.GET("GetUserInfo", h.UserApi.GetUserInfo) // 获取自身信息

		RouterCasbinUserApi.POST("GetUserList", h.UserApi.GetUserList)               // 分页获取用户列表
		RouterCasbinUserApi.POST("SetUserAuthorities", h.UserApi.SetUserAuthorities) // 设置用户权限组
		RouterCasbinUserApi.POST("ResetPassword", h.UserApi.ResetPassword)           // 设置用户权限组
		RouterCasbinUserApi.POST("ChangePassword", h.UserApi.ChangePassword)         // 用户修改密码
		RouterCasbinUserApi.POST("SetUserAuthority", h.UserApi.SetUserAuthority)     // 设置用户权限
		RouterCasbinUserApi.DELETE("DeleteUser", h.UserApi.DeleteUser)               // 删除用户
		RouterCasbinUserApi.PUT("SetUserInfo", h.UserApi.SetUserInfo)                // 设置用户信息
	}

	//权限控制模块
	RouterCasbinApi := r.Group(service.ImplMap["sys_casbin"])
	{
		RouterCasbinApi.POST("GetPolicyPathByAuthorityId", h.CasbinApi.GetPolicyPathByAuthorityId)
		RouterCasbinApi.POST("UpdateCasbin", h.CasbinApi.UpdateCasbin)
	}
}

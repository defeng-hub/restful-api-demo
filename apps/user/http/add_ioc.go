package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps"
	"restful-api-demo/apps/user"
	"restful-api-demo/apps/user/service"
)

type Handler struct {
	sys_user *service.UserService
	sys_menu *service.MenuService
}

func (h *Handler) Name() string {
	return user.AppName
}

// Registry 注册路由
func (h *Handler) Registry(r gin.IRouter) {
	RouterUserService := r.Group(service.ImplMap["sys_user"])
	{
		RouterUserService.GET("aaa")
	}

	RouterMenuService := r.Group(service.ImplMap["sys_menu"])
	{
		RouterMenuService.POST("aaa")

	}
}

func (h *Handler) Config() {
	h.sys_user = apps.GetImpl(new(service.UserService).Name()).(*service.UserService)
	h.sys_menu = apps.GetImpl(new(service.MenuService).Name()).(*service.MenuService)
}

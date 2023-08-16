package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps"
	"restful-api-demo/apps/user"
	"restful-api-demo/apps/user/http/api"
	"restful-api-demo/apps/user/service"
)

type Handler struct {
	//UserApi   api.UserApi
	CasbinApi api.CasbinApi
}

func (h *Handler) Name() string {
	return user.AppName
}

func (h *Handler) Config() {
	//h.UserApi.Srv = apps.GetImpl(new(service.UserService).Name()).(*service.UserService)
	h.CasbinApi.Srv = apps.GetImpl(new(service.CasbinService).Name()).(*service.CasbinService)
}

// Registry 注册路由
func (h *Handler) Registry(r gin.IRouter) {
	Router(h, r)
}

package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps"
	"restful-api-demo/apps/user"
	"restful-api-demo/apps/user/http/api"
	. "restful-api-demo/apps/user/service"
	"restful-api-demo/common/logger/zap"
)

type Handler struct {
	UserApi   api.UserApi
	CasbinApi api.CasbinApi
}

func (h *Handler) Name() string {
	return user.AppName
}

func (h *Handler) Config() {
	// 对各个API的  内部实现类进行初始化
	// TODO:需要在这注册每一个app
	h.UserApi = api.UserApi{
		Srv: apps.GetImpl(new(UserService).Name()).(*UserService),
		L:   zap.L().Named(new(UserService).Name()),
	}

	h.CasbinApi.Srv = apps.GetImpl(new(CasbinService).Name()).(*CasbinService)
}

// Registry 注册路由
func (h *Handler) Registry(r gin.IRouter) {
	Router(h, r)
}

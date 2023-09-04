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
	UserApi      api.UserApi
	CasbinApi    api.CasbinApi
	MenuApi      api.AuthorityMenuApi
	AuthorityApi api.AuthorityApi
	SystemApiApi api.SystemApiApi
}

func (h *Handler) Name() string {
	return user.AppName
}

func (h *Handler) Config() {
	// 对各个API的  内部实现类进行初始化
	h.UserApi = api.UserApi{
		Srv: apps.GetImpl(new(UserService).Name()).(*UserService),
		L:   zap.L().Named(new(UserService).Name()),
	}

	h.CasbinApi.Srv = apps.GetImpl(new(CasbinService).Name()).(*CasbinService)

	h.MenuApi = api.AuthorityMenuApi{
		Srv: apps.GetImpl(new(MenuService).Name()).(*MenuService),
		L:   zap.L().Named(new(MenuService).Name()),
	}

	h.AuthorityApi = api.AuthorityApi{
		Srv: apps.GetImpl(new(AuthorityService).Name()).(*AuthorityService),
		L:   zap.L().Named(new(AuthorityService).Name()),
	}
	h.SystemApiApi = api.SystemApiApi{
		Srv: apps.GetImpl(new(ApiService).Name()).(*ApiService),
		L:   zap.L().Named(new(ApiService).Name()),
	}
}

// Registry 注册路由
func (h *Handler) Registry(r gin.IRouter) {
	Router(h, r)
}

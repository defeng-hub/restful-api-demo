package http

import (
	"github.com/defeng-hub/restful-api-demo/apps"
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc host.Service
}

// 用来注册到ioc
var handler = &Handler{}

func (h *Handler) Name() string {
	return host.AppName
}

// Registry 注册路由
func (h *Handler) Registry(r gin.IRouter) {
	r.GET("/"+host.AppName+"/hosts", h.queryHost)
	r.GET("/"+host.AppName+"/hosts/:id", h.describeHost)
	r.POST("/"+host.AppName+"/hosts", h.createHost)
	r.PUT("/"+host.AppName+"/hosts/:id", h.putHost)
	r.PATCH("/"+host.AppName+"/hosts/:id", h.patchHost)
}

func (h *Handler) Config() {
	h.svc = apps.GetImpl(host.AppName).(host.Service)
	if h.svc == nil {
		panic("在IOC中 没有获取到HostService")
	}
}

// 因为实现了上述的三个函数   所以可以注册进去ioc
func init() {
	apps.RegistryGin(handler)
}

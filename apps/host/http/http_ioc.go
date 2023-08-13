package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps"
	"restful-api-demo/apps/host"
)

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

package http

import (
	"github.com/defeng-hub/restful-api-demo/apps"
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc host.Service
}

var handler = &Handler{}

func (h *Handler) Name() string {
	return host.AppName
}

// Registry http handler注册
func (h *Handler) Registry(r gin.IRouter) {
	r.GET("/"+host.AppName+"/hosts", h.createHost)
	r.POST("/"+host.AppName+"/hosts", h.createHost)
}

func (h *Handler) Config() {
	h.svc = apps.GetImpl(host.AppName).(host.Service)
	if h.svc == nil {
		panic("在IOC中 没有获取到HostService")
	}
}
func init() {
	apps.RegistryGin(handler)
}

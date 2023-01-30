package http

import (
	"github.com/defeng-hub/restful-api-demo/apps"
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc host.Service // 选择接口就可以采用mysql实现获取其他类型实现
}

var handler = &Handler{}

// NewHttpHandler 面向接口, 真正service的实现, 在服务实例化的时候传进来
func NewHttpHandler() *Handler {
	return &Handler{}
}

// Registry http handler注册
func (h *Handler) Registry(r gin.IRouter) {
	r.GET("/hosts", h.createHost)
	r.POST("/hosts", h.createHost)
}

func (h *Handler) Config() {
	h.svc = apps.GetImpl(host.AppName).(host.Service)
	if h.svc == nil {
		panic("在IOC中 没有获取到HostService")
	}

}

func (h *Handler) Name() string {
	return host.AppName
}

func init() {
	apps.RegistryGin(handler)
}

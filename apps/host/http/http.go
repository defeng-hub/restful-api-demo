package http

import (
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
)

// Handler 暴露出去接口,
type Handler struct {
	svc host.Service // 选择接口就可以采用mysql实现获取其他类型实现
}

func NewHttpHandler(svc host.Service) *Handler {
	return &Handler{svc: svc}
}

// Registry http handler注册
func (h *Handler) Registry(r gin.IRouter) {
	r.GET("/hosts", h.createHost)
	r.POST("/hosts", h.createHost)
}

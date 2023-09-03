package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/system"
)

type Handler struct{}

func (h *Handler) Name() string {
	return system.AppName
}

func (h *Handler) Config() {}

// Registry 注册路由
func (h *Handler) Registry(r gin.IRouter) {
	Router(h, r)
}

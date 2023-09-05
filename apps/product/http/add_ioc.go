package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps"
	"restful-api-demo/apps/product"
	"restful-api-demo/apps/product/impl"
)

func (h *Handler) Name() string {
	return product.AppName
}

func (h *Handler) Config() {
	h.proSrv = apps.GetImpl(new(impl.ProductImpl).Name()).(product.ProductInterface)
}

// Registry 注册路由
func (h *Handler) Registry(r gin.IRouter) {
	Router(h, r)
}

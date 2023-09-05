package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/user/http/middleware"
)

var Private = []gin.HandlerFunc{middleware.JWTAuth(), middleware.CasbinHandler()}
var Public = []gin.HandlerFunc{}

func Router(h *Handler, r gin.IRouter) {
	//产品模块
	//RouterProductApi := r.Group("UserService").Use(Public...)
	RouterCasbinProductApi := r.Group("ProductService").Use(Public...)
	{
		RouterCasbinProductApi.POST("create", h.createProduct)
		RouterCasbinProductApi.POST("query", h.queryProduct)
	}
}

package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/user/http/middleware"
)

var Private = []gin.HandlerFunc{middleware.JWTAuth(), middleware.CasbinHandler(), middleware.OperationRecord()}
var Public = []gin.HandlerFunc{}

func Router(h *Handler, r gin.IRouter) {
	Api := r.Group("SystemService").Use(Public...)
	{
		Api.POST("Info", h.GetSystemInfo)
	}
}

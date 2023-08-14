package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/user/service"
)

func Router(h *Handler, r gin.IRouter) {
	RouterUserApi := r.Group(service.ImplMap["sys_user"])
	{
		RouterUserApi.GET("aaa")
	}
	RouterCasbinApi := r.Group(service.ImplMap["sys_casbin"])
	{
		RouterCasbinApi.POST("getPolicyPathByAuthorityId", h.CasbinApi.GetPolicyPathByAuthorityId)
	}
}

package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/user/common/response"
	"restful-api-demo/conf"
)

func (h *Handler) GetSystemInfo(c *gin.Context) {
	var SysConfigResponse struct {
		Config *conf.Config `json:"config" comment:"系统配置"`
	}
	SysConfigResponse.Config = conf.C()
	response.OkWithDetailed(SysConfigResponse, "获取成功", c)
}

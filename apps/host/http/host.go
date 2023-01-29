package http

import (
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

//创建host的http-api
func (h *Handler) createHost(c *gin.Context) {
	//获取用户传递的参数,并解析
	ins := &host.Host{
		Resource: &host.Resource{},
		Describe: &host.Describe{},
	}
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}

	resIns, err := h.svc.SaveHost(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}
	response.Success(c.Writer, resIns)
}

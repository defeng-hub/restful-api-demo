package http

import (
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
)

type Handler struct {
	svc host.Service // 选择接口就可以采用mysql实现获取其他类型实现
}

//创建host的http-api
func (h *Handler) createHost(c *gin.Context) {
	//获取用户传递的参数,并解析
	ins := host.NewHost()
	if err := c.Bind(ins); err != nil {
		response.Failed(c.Writer, err)
		return
	}
	//校验参数
	err := ins.Validate()
	if err != nil {
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

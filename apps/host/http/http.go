package http

import (
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
	"strconv"
)

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

func (h *Handler) queryHost(c *gin.Context) {
	req := host.NewQueryHostRequest(0, 0, "")
	query := c.Request.URL.Query()
	pss := query.Get("page_size")
	pns := query.Get("page_number")
	if pns != "" {
		ls, _ := strconv.Atoi(pns)
		req.PageNumber = uint64(ls)
	}
	if pss != "" {
		ls, _ := strconv.Atoi(pss)
		req.PageSize = uint64(ls)
	}

	set, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		return
	}
	response.Success(c.Writer, set)
}

package http

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/product"
	"restful-api-demo/apps/product/common/response"
	"restful-api-demo/apps/product/model"
)

type Handler struct {
	proSrv product.ProductInterface
}

func (h *Handler) createProduct(c *gin.Context) {
	//获取用户传递的参数,并解析
	pro := model.NewProduct("产品名")
	if err := c.Bind(pro); err != nil {
		response.Fail(c, err)
		return
	}

	//校验参数
	err := pro.Validate()
	if err != nil {
		response.Fail(c, err)
		return
	}

	resPro, err := h.proSrv.SaveProduct(c.Request.Context(), pro)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.OkWithDetailed(resPro, "添加成功", c)
}

func (h *Handler) queryProduct(c *gin.Context) {
	req := &model.QueryProductListRequest{}
	err := c.Bind(req)
	if err != nil {
		response.Fail(c, err)
		return
	}

	proSet, err := h.proSrv.QueryProduct(c.Request.Context(), req)
	if err != nil {
		return
	}
	response.OkWithDetailed(proSet, "获取成功", c)
}

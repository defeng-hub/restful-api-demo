package api

import (
	"github.com/gin-gonic/gin"
	"restful-api-demo/apps/user/common/response"
	"restful-api-demo/apps/user/model/request"
	systemRes "restful-api-demo/apps/user/model/response"
	"restful-api-demo/apps/user/service"
)

type CasbinApi struct {
	Srv *service.CasbinService
}

// UpdateCasbin
// @Summary 更新角色api权限
func (cas *CasbinApi) UpdateCasbin(c *gin.Context) {
	var cmr request.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)

	if err := cas.Srv.UpdateCasbin(cmr.AuthorityId, cmr.CasbinInfos); err != nil {
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetPolicyPathByAuthorityId
// @Summary 获取权限列表
func (cas *CasbinApi) GetPolicyPathByAuthorityId(c *gin.Context) {
	var casbin request.CasbinInReceive
	_ = c.ShouldBindJSON(&casbin)
	if casbin.AuthorityId == "" {
		response.FailWithMessage("未输入权限id", c)
		return
	}
	paths := cas.Srv.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "获取成功", c)
	return
}

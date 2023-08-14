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
// @Tags Casbin
// @Summary 更新角色api权限
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Response{msg=string} "更新角色api权限"
// @Router /casbin/UpdateCasbin [post]
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
// @Tags Casbin
// @Summary 获取权限列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CasbinInReceive true "权限id, 权限模型列表"
// @Success 200 {object} response.Response{data=systemRes.PolicyPathResponse,msg=string} "获取权限列表,返回包括casbin详情列表"
// @Router /casbin/getPolicyPathByAuthorityId [post]
func (cas *CasbinApi) GetPolicyPathByAuthorityId(c *gin.Context) {
	var casbin request.CasbinInReceive
	_ = c.ShouldBindJSON(&casbin)
	paths := cas.Srv.GetPolicyPathByAuthorityId(casbin.AuthorityId)
	response.OkWithDetailed(systemRes.PolicyPathResponse{Paths: paths}, "获取成功", c)
}

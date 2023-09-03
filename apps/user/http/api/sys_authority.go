package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	utils "restful-api-demo/apps/user/common"
	"restful-api-demo/apps/user/common/request"
	"restful-api-demo/apps/user/common/response"
	"restful-api-demo/apps/user/model"
	systemRes "restful-api-demo/apps/user/model/response"
	"restful-api-demo/apps/user/service"
	"restful-api-demo/common/logger"
)

type AuthorityApi struct {
	Srv *service.AuthorityService
	L   logger.Logger
}

// @Summary 创建角色
func (a *AuthorityApi) CreateAuthority(c *gin.Context) {
	var authority *model.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := a.Srv.CreateAuthority(authority); err != nil {
		a.L.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: *authBack}, "创建成功", c)
	}
}

// @Summary 拷贝角色
func (a *AuthorityApi) CopyAuthority(c *gin.Context) {
	var copyInfo systemRes.SysAuthorityCopyResponse
	_ = c.ShouldBindJSON(&copyInfo)
	if err := utils.Verify(copyInfo, utils.OldAuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(copyInfo.Authority, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authBack := a.Srv.CopyAuthority(copyInfo); err != nil {
		a.L.Error("拷贝失败!", zap.Error(err))
		response.FailWithMessage("拷贝失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authBack}, "拷贝成功", c)
	}
}

// @Summary 删除角色
func (a *AuthorityApi) DeleteAuthority(c *gin.Context) {
	var authority model.SysAuthority
	_ = c.ShouldBindJSON(&authority)
	if err := utils.Verify(authority, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := a.Srv.DeleteAuthority(&authority); err != nil { // 删除角色之前需要判断是否有用户正在使用此角色
		a.L.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败"+err.Error(), c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Summary 更新角色信息
func (a *AuthorityApi) UpdateAuthority(c *gin.Context) {
	var auth model.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, authority := a.Srv.UpdateAuthority(auth); err != nil {
		a.L.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(systemRes.SysAuthorityResponse{Authority: authority}, "更新成功", c)
	}
}

// @Summary 分页获取角色列表
func (a *AuthorityApi) GetAuthorityList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := a.Srv.GetAuthorityInfoList(pageInfo); err != nil {
		a.L.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Summary 设置角色资源权限
func (a *AuthorityApi) SetDataAuthority(c *gin.Context) {
	var auth model.SysAuthority
	_ = c.ShouldBindJSON(&auth)
	if err := utils.Verify(auth, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := a.Srv.SetDataAuthority(auth); err != nil {
		a.L.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败"+err.Error(), c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

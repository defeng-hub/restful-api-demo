package api

import (
	utils "restful-api-demo/apps/user/common"
	"restful-api-demo/apps/user/common/request"
	"restful-api-demo/apps/user/common/response"
	"restful-api-demo/apps/user/model"
	"restful-api-demo/apps/user/service"
	"restful-api-demo/common/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemApiApi struct {
	Srv *service.ApiService
	L   logger.Logger
}

// @Tags SysApi
// @Summary 创建基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {object} response.Response{msg=string} "创建基础api"
// @Router /api/createApi [post]
func (s *SystemApiApi) CreateApi(c *gin.Context) {
	var api model.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := s.Srv.CreateApi(api); err != nil {
		s.L.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SysApi
// @Summary 删除api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "ID"
// @Success 200 {object} response.Response{msg=string} "删除api"
// @Router /api/deleteApi [post]
func (s *SystemApiApi) DeleteApi(c *gin.Context) {
	var api model.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api.BASEMODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := s.Srv.DeleteApi(api); err != nil {
		s.L.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SysApi
// @Summary 分页获取API列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SearchApiParams true "分页获取API列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取API列表,返回包括列表,总数,页码,每页数量"
// @Router /api/getApiList [post]
func (s *SystemApiApi) GetApiList(c *gin.Context) {
	// api分页条件查询及排序结构体
	type SearchApiParams struct {
		model.SysApi
		request.PageInfo
		OrderKey string `json:"orderKey"` // 排序
		Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
	}
	var pageInfo SearchApiParams
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := s.Srv.GetAPIInfoList(pageInfo.SysApi, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc); err != nil {
		s.L.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// todo
// @Tags SysApi
// @Summary 根据id获取api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "根据id获取api"
// @Success 200 {object} response.Response{data=systemRes.SysAPIResponse} "根据id获取api,返回包括api详情"
// @Router /api/getApiById [post]
func (s *SystemApiApi) GetApiById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, api := s.Srv.GetApiById(idInfo.ID)
	if err != nil {
		s.L.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		var SysAPIResponse struct {
			Api model.SysApi `json:"api" comment:"api详情"`
		}
		SysAPIResponse.Api = api
		response.OkWithData(SysAPIResponse, c)
	}
}

// @Tags SysApi
// @Summary 修改基础api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SysApi true "api路径, api中文描述, api组, 方法"
// @Success 200 {object} response.Response{msg=string} "修改基础api"
// @Router /api/updateApi [post]
func (s *SystemApiApi) UpdateApi(c *gin.Context) {
	var api model.SysApi
	_ = c.ShouldBindJSON(&api)
	if err := utils.Verify(api, utils.ApiVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := s.Srv.UpdateApi(api); err != nil {
		s.L.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

// @Tags SysApi
// @Summary 获取所有的Api 不分页
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=systemRes.SysAPIListResponse,msg=string} "获取所有的Api 不分页,返回包括api列表"
// @Router /api/getAllApis [post]
func (s *SystemApiApi) GetAllApis(c *gin.Context) {
	if err, apis := s.Srv.GetAllApis(); err != nil {
		s.L.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		var SysAPIListResponse struct {
			Apis []model.SysApi `json:"apis" comment:"api列表"`
		}
		SysAPIListResponse.Apis = apis
		response.OkWithDetailed(SysAPIListResponse, "获取成功", c)
	}
}

// @Tags SysApi
// @Summary 删除选中Api
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "ID"
// @Success 200 {object} response.Response{msg=string} "删除选中Api"
// @Router /api/deleteApisByIds [delete]
func (s *SystemApiApi) DeleteApisByIds(c *gin.Context) {
	var ids request.IdsReq
	_ = c.ShouldBindJSON(&ids)
	if err := s.Srv.DeleteApisByIds(ids); err != nil {
		s.L.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

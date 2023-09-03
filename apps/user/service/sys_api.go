package service

import (
	"errors"
	"restful-api-demo/apps"
	"restful-api-demo/apps/user"
	"restful-api-demo/apps/user/common/request"
	"restful-api-demo/apps/user/model"
	"restful-api-demo/conf"

	"gorm.io/gorm"
)

type ApiService struct {
	db        *gorm.DB
	casbinSrv *CasbinService
}

func (s *ApiService) Name() string {
	return user.AppName + ImplMap["sys_api"]
}
func (s *ApiService) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
	s.casbinSrv = apps.GetImpl(new(CasbinService).Name()).(*CasbinService)
}

//@description: 新增基础api
func (apiService *ApiService) CreateApi(api model.SysApi) (err error) {
	if !errors.Is(apiService.db.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return apiService.db.Create(&api).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApi(api model.SysApi) (err error) {
	err = apiService.db.Delete(&api).Error
	apiService.casbinSrv.ClearCasbin(1, api.Path, api.Method)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error

func (apiService *ApiService) GetAPIInfoList(api model.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := apiService.db.Model(&model.SysApi{})
	var apiList []model.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			// 感谢 Tom4t0 提交漏洞信息
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			}

			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi

func (apiService *ApiService) GetAllApis() (err error, apis []model.SysApi) {
	err = apiService.db.Find(&apis).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi

func (apiService *ApiService) GetApiById(id float64) (err error, api model.SysApi) {
	err = apiService.db.Where("id = ?", id).First(&api).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) UpdateApi(api model.SysApi) (err error) {
	var oldA model.SysApi
	err = apiService.db.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(apiService.db.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = apiService.casbinSrv.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = apiService.db.Save(&api).Error
		}
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApis
//@description: 删除选中API
//@param: apis []model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	err = apiService.db.Delete(&[]model.SysApi{}, "id in ?", ids.Ids).Error
	return err
}

func (apiService *ApiService) DeleteApiByIds(ids []string) (err error) {
	return apiService.db.Delete(&model.SysApi{}, "id in ?", ids).Error
}

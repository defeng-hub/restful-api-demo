package service

import (
	"errors"
	"gorm.io/gorm"
	"restful-api-demo/apps"
	"restful-api-demo/apps/user"
	"restful-api-demo/apps/user/common/request"
	"restful-api-demo/apps/user/model"
	"restful-api-demo/apps/user/model/response"
	"restful-api-demo/conf"
	"strconv"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAuthority
//@description: 创建一个角色
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority

type AuthorityService struct {
	db        *gorm.DB
	menuSrv   *MenuService
	casbinSrv *CasbinService
}

func (s *AuthorityService) Name() string {
	return user.AppName + ImplMap["sys_authority"]
}
func (s *AuthorityService) Config() { // base_sys_menu.go
	s.db, _ = conf.C().MySQL.GetGormDB()
	s.menuSrv = apps.GetImpl(new(MenuService).Name()).(*MenuService)
	s.casbinSrv = apps.GetImpl(new(CasbinService).Name()).(*CasbinService)
}

func (a *AuthorityService) CreateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	var authorityBox model.SysAuthority
	if !errors.Is(a.db.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = a.db.Create(&auth).Error
	return err, auth
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CopyAuthority
//@description: 复制一个角色
//@param: copyInfo response.SysAuthorityCopyResponse
//@return: err error, authority model.SysAuthority

func (a *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (err error, authority model.SysAuthority) {
	var authorityBox model.SysAuthority
	if !errors.Is(a.db.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), authority
	}
	copyInfo.Authority.Children = []model.SysAuthority{}
	err, menus := a.menuSrv.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	if err != nil {
		return
	}
	var baseMenu []model.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = a.db.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}
	paths := a.casbinSrv.GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = a.casbinSrv.UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = a.DeleteAuthority(&copyInfo.Authority)
	}
	return err, copyInfo.Authority
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAuthority
//@description: 更改一个角色
//@param: auth model.SysAuthority
//@return: err error, authority model.SysAuthority

func (a *AuthorityService) UpdateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	err = a.db.Where("authority_id = ?", auth.AuthorityId).First(&model.SysAuthority{}).Updates(&auth).Error
	return err, auth
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAuthority
//@description: 删除角色
//@param: auth *model.SysAuthority
//@return: err error

func (a *AuthorityService) DeleteAuthority(auth *model.SysAuthority) (err error) {
	if errors.Is(a.db.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(a.db.Where("authority_id = ?", auth.AuthorityId).First(&model.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(a.db.Where("parent_id = ?", auth.AuthorityId).First(&model.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := a.db.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if err != nil {
		return
	}
	if len(auth.SysBaseMenus) > 0 {
		err = a.db.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		if err != nil {
			return
		}
		// err = db.Association("SysBaseMenus").Delete(&auth)
	} else {
		err = db.Error
		if err != nil {
			return
		}
	}
	err = a.db.Delete(&[]model.SysUseAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error

	//清除 casbin权限
	a.casbinSrv.ClearCasbin(0, auth.AuthorityId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfoList
//@description: 分页获取数据
//@param: info request.PageInfo
//@return: err error, list interface{}, total int64

func (a *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := a.db.Model(&model.SysAuthority{})
	err = db.Where("parent_id = ?", "0").Count(&total).Error
	var authority []model.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = ?", "0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = a.findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorityInfo
//@description: 获取所有角色信息
//@param: auth model.SysAuthority
//@return: err error, sa model.SysAuthority

func (a *AuthorityService) GetAuthorityInfo(auth model.SysAuthority) (err error, sa model.SysAuthority) {
	err = a.db.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetDataAuthority
//@description: 设置角色资源权限
//@param: auth model.SysAuthority
//@return: error

func (a *AuthorityService) SetDataAuthority(auth model.SysAuthority) error {
	var s model.SysAuthority
	a.db.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := a.db.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetMenuAuthority
//@description: 菜单与角色绑定
//@param: auth *model.SysAuthority
//@return: error

func (a *AuthorityService) SetMenuAuthority(auth *model.SysAuthority) error {
	var s model.SysAuthority
	a.db.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := a.db.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: findChildrenAuthority
//@description: 查询子角色
//@param: authority *model.SysAuthority
//@return: err error

func (a *AuthorityService) findChildrenAuthority(authority *model.SysAuthority) (err error) {
	err = a.db.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = a.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

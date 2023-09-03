package service

import (
	"errors"
	"restful-api-demo/apps"
	"restful-api-demo/apps/user"
	"restful-api-demo/conf"
	"strconv"

	"restful-api-demo/apps/user/common/request"
	"restful-api-demo/apps/user/model"

	"gorm.io/gorm"
)

//@function: getMenuTreeMap
//@description: 获取路由总树map
//@param: authorityId string
//@return: err error, treeMap map[string][]model.SysMenu

type MenuService struct {
	db      *gorm.DB
	authSrv *AuthorityService
}

func (s *MenuService) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
	s.authSrv = apps.GetImpl(new(AuthorityService).Name()).(*AuthorityService)
}
func (s *MenuService) Name() string {
	return user.AppName + ImplMap["sys_menu"]
}

func (menuService *MenuService) getMenuTreeMap(authorityId string) (err error, treeMap map[string][]model.SysMenu) {
	var allMenus []model.SysMenu
	treeMap = make(map[string][]model.SysMenu)
	err = menuService.db.Where("authority_id = ?", authorityId).Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@description: 获取动态菜单树
func (menuService *MenuService) GetMenuTree(authorityId string) (err error, menus []model.SysMenu) {
	err, menuTree := menuService.getMenuTreeMap(authorityId)
	menus = menuTree["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return err, menus
}

//@description: 获取子菜单
//@return: err error
func (menuService *MenuService) getChildrenList(menu *model.SysMenu, treeMap map[string][]model.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuId]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@description: 获取路由分页
//@return: err error, list interface{}, total int64
func (menuService *MenuService) GetInfoList() (err error, list interface{}, total int64) {
	var menuList []model.SysBaseMenu
	err, treeMap := menuService.getBaseMenuTreeMap()
	menuList = treeMap["0"]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return err, menuList, total
}

//@description: 获取菜单的子菜单
//@return: err error
func (menuService *MenuService) getBaseChildrenList(menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu) (err error) {
	menu.Children = treeMap[strconv.Itoa(int(menu.ID))]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

//@description: 添加基础路由
//@return: error
func (menuService *MenuService) AddBaseMenu(menu model.SysBaseMenu) error {
	if !errors.Is(menuService.db.Where("name = ?", menu.Name).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return menuService.db.Create(&menu).Error
}

//@description: 获取路由总树map
//@return: err error, treeMap map[string][]model.SysBaseMenu
func (menuService *MenuService) getBaseMenuTreeMap() (err error, treeMap map[string][]model.SysBaseMenu) {
	var allMenus []model.SysBaseMenu
	treeMap = make(map[string][]model.SysBaseMenu)
	err = menuService.db.Order("sort").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return err, treeMap
}

//@description: 获取基础路由树
//@return: err error, menus []model.SysBaseMenu
func (menuService *MenuService) GetBaseMenuTree() (err error, menus []model.SysBaseMenu) {
	err, treeMap := menuService.getBaseMenuTreeMap()
	menus = treeMap["0"]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return err, menus
}

//@description: 为角色增加menu树
//@return: err error
func (menuService *MenuService) AddMenuAuthority(menus []model.SysBaseMenu, authorityId string) (err error) {
	var auth model.SysAuthority
	auth.AuthorityId = authorityId
	auth.SysBaseMenus = menus

	err = menuService.authSrv.SetMenuAuthority(&auth)
	return err
}

//@description: 查看指定角色树
func (menuService *MenuService) GetMenuAuthority(info *request.GetAuthorityId) (err error, menus []model.SysMenu) {
	err = menuService.db.Where("authority_id = ? ", info.AuthorityId).Order("sort").Find(&menus).Error
	// sql := "SELECT authority_menu.keep_alive,authority_menu.default_menu,authority_menu.created_at,authority_menu.updated_at,authority_menu.deleted_at,authority_menu.menu_level,authority_menu.parent_id,authority_menu.path,authority_menu.`name`,authority_menu.hidden,authority_menu.component,authority_menu.title,authority_menu.icon,authority_menu.sort,authority_menu.menu_id,authority_menu.authority_id FROM authority_menu WHERE authority_menu.authority_id = ? ORDER BY authority_menu.sort ASC"
	// err = menuService.db.Raw(sql, authorityId).Scan(&menus).Error
	return err, menus
}

// 基础路由，基础路由，基础路由，基础路由

//@description: 删除基础路由
func (baseMenuService *MenuService) DeleteBaseMenu(id float64) (err error) {
	err = baseMenuService.db.Preload("Parameters").Where("parent_id = ?", id).First(&model.SysBaseMenu{}).Error
	if err != nil {
		var menu model.SysBaseMenu
		db := baseMenuService.db.Preload("SysAuthoritys").Where("id = ?", id).First(&menu).Delete(&menu)
		err = baseMenuService.db.Delete(&model.SysBaseMenuParameter{}, "sys_base_menu_id = ?", id).Error
		if err != nil {
			return err
		}
		if len(menu.SysAuthoritys) > 0 {
			err = baseMenuService.db.Model(&menu).Association("SysAuthoritys").Delete(&menu.SysAuthoritys)
		} else {
			err = db.Error
			if err != nil {
				return
			}
		}
	} else {
		return errors.New("此菜单存在子菜单不可删除")
	}
	return err
}

//@description: 更新路由
func (baseMenuService *MenuService) UpdateBaseMenu(menu model.SysBaseMenu) (err error) {
	var oldMenu model.SysBaseMenu
	upDateMap := make(map[string]interface{})
	upDateMap["keep_alive"] = menu.KeepAlive
	upDateMap["close_tab"] = menu.CloseTab
	upDateMap["default_menu"] = menu.DefaultMenu
	upDateMap["parent_id"] = menu.ParentId
	upDateMap["path"] = menu.Path
	upDateMap["name"] = menu.Name
	upDateMap["hidden"] = menu.Hidden
	upDateMap["component"] = menu.Component
	upDateMap["title"] = menu.Title
	upDateMap["icon"] = menu.Icon
	upDateMap["sort"] = menu.Sort

	err = baseMenuService.db.Transaction(func(tx *gorm.DB) error {
		db := tx.Where("id = ?", menu.ID).Find(&oldMenu)
		if oldMenu.Name != menu.Name {
			if !errors.Is(tx.Where("id <> ? AND name = ?", menu.ID, menu.Name).First(&model.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
				return errors.New("存在相同name修改失败")
			}
		}
		txErr := tx.Unscoped().Delete(&model.SysBaseMenuParameter{}, "sys_base_menu_id = ?", menu.ID).Error
		if txErr != nil {
			return txErr
		}
		if len(menu.Parameters) > 0 {
			for k := range menu.Parameters {
				menu.Parameters[k].SysBaseMenuID = menu.ID
			}
			txErr = tx.Create(&menu.Parameters).Error
			if txErr != nil {
				return txErr
			}
		}

		txErr = db.Updates(upDateMap).Error
		if txErr != nil {
			return txErr
		}
		return nil
	})
	return err
}

//@description: 返回当前选中menu
func (baseMenuService *MenuService) GetBaseMenuById(id float64) (err error, menu model.SysBaseMenu) {
	err = baseMenuService.db.Preload("Parameters").Where("id = ?", id).First(&menu).Error
	return
}

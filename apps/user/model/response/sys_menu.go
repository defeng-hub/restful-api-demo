package response

import "restful-api-demo/apps/user/model"

type SysMenusResponse struct {
	Menus []model.SysMenu `json:"menus" comment:"系统菜单详情列表"`
}

type SysBaseMenusResponse struct {
	Menus []model.SysBaseMenu `json:"menus" comment:"系统菜单列表"`
}

type SysBaseMenuResponse struct {
	Menu model.SysBaseMenu `json:"menu" comment:"系统菜单列表"`
}

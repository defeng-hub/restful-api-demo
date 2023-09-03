package model

type SysApi struct {
	BASEMODEL
	Path        string `json:"path" gorm:"comment:api路径" example:"/api"`            // api路径
	Description string `json:"description" gorm:"comment:api中文描述" example:"测试api"`  // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组" example:"api_group"`    // api组
	Method      string `json:"method" gorm:"default:POST;comment:方法" example:"GET"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

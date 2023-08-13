package model

import (
	"gorm.io/gorm"
	"time"
)

type SysAuthority struct {
	CreatedAt       time.Time      // 创建时间
	UpdatedAt       time.Time      // 更新时间
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`                                                                                                     // 删除时间
	AuthorityId     string         `json:"authorityId" gorm:"primary_key;not null;unique;comment:角色ID;size:90" example:"e07122e1-56ee-4016-8070-2d9d3e6550ab"` // 角色ID
	AuthorityName   string         `json:"authorityName" gorm:"comment:角色名" example:"admin2"`                                                                  // 角色名
	ParentId        string         `json:"parentId" gorm:"comment:父角色ID" example:"admin"`                                                                      // 父角色ID
	DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu  `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users           []SysUser      `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter   string         `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

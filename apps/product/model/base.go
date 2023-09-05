package model

import (
	"time"

	"gorm.io/gorm"
)

type UPDATE_MODE string

const (
	UPDATE_MODE_PUT   UPDATE_MODE = "put"   // 全量更新
	UPDATE_MODE_PATCH UPDATE_MODE = "patch" // 局部更新
)

type BASEMODEL struct {
	ID        uint           `gorm:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type Page struct {
	OrderBy  string `json:"order_by"`
	Desc     bool   `json:"desc"`
	Page     int64  `json:"page"`
	Size     int64  `json:"size"`
	Keywords string `json:"kws"`
}

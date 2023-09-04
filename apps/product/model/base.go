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
	OrderBy string `json:"orderBy"`
	Asc     bool   `json:"asc"`
	Page    int    `json:"page"`
	Size    int    `json:"size"`
}

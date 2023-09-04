package model

import (
	"dario.cat/mergo"
	"fmt"
	"github.com/go-playground/validator/v10"
)

var (
	// validate 必须先new一个实例，然后在实例上使用 v.Struct( obj )，  obj 就是 某个结构体的实例
	validate = validator.New()
)

type Product struct {
	BASEMODEL
	Name  string
	Price float64
	Count int64
}

func (p *Product) TableName() string {
	return "p_product"
}

// Validate  Product参数校验
func (p *Product) Validate() error {
	return validate.Struct(p)
}

// 对象全量更新
func (p *Product) Put(obj *Product) error {
	if obj.ID != p.ID {
		return fmt.Errorf("id not equal")
	}
	*p = *obj
	return nil
}

// 对象的局部更新
func (p *Product) Patch(obj *Product) error {
	// 比如 h.A  obj.B  只想修改obj.B该属性
	//return mergo.MergeWithOverwrite(h, obj)
	return mergo.Merge(p, *obj, mergo.WithOverride)
}

// ----------------------------
type QueryProductListRequest struct {
	OrderBy  string `json:"order_by"`
	Asc      bool   `json:"asc"`
	Page     int    `json:"page"`
	Size     int    `json:"size"`
	Keywords string `json:"kws"`
}
type DescribeProductRequest struct {
	Id string `json:"id"`
}
type DeleteProductRequest struct {
	Id string `json:"id"`
}
type UpdateProductRequest struct {
	UpdateMode UPDATE_MODE `json:"update_mode"`
	*Product
}
type ProductSet struct {
	Total int        `json:"total"`
	Items []*Product `json:"items"`
}

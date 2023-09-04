package product

import (
	"context"
	"restful-api-demo/apps/product/model"
)

type ProductInterface interface {
	SaveProduct(context.Context, *model.Product) (*model.Product, error)
	QueryProduct(context.Context, *model.QueryProductListRequest) (*model.ProductSet, error)
	DescribeProduct(context.Context, *model.DescribeProductRequest) (*model.Product, error)
	DeleteProduct(context.Context, *model.DeleteProductRequest) (*model.Product, error)
	UpdateProduct(context.Context, *model.UpdateProductRequest) (*model.Product, error)
}

//type Coupon interface {
//	InsertCoupon(pro *Coupon) *Coupon
//	DeleteCoupon(pro *Coupon) *Coupon
//	UpdateCoupon(pro *Coupon) *Coupon
//	GetCouponList(p *Page) []*ProductSet
//}

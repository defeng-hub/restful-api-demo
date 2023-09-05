package impl

import (
	"context"
	"gorm.io/gorm"
	"restful-api-demo/apps/product"
	"restful-api-demo/apps/product/model"
	"restful-api-demo/common/logger"
	"restful-api-demo/conf"
	"restful-api-demo/utils/sqlbuilder"
)

type ProductImpl struct {
	db *gorm.DB
	l  logger.Logger
}

func (s *ProductImpl) Name() string {
	return product.AppName + "(" + productService + ")"
}

func (s *ProductImpl) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
	s.l = conf.L().Named(productService)
}

//编译器做静态检查
var _ product.ProductInterface = (*ProductImpl)(nil)

func (p *ProductImpl) SaveProduct(ctx context.Context, m *model.Product) (*model.Product, error) {
	// 注入默认值
	m.InjectDefault()

	err := m.Validate()
	if err != nil {
		return nil, err
	}
	return m, p.db.Create(m).Error
}

func (p *ProductImpl) QueryProduct(ctx context.Context, req *model.QueryProductListRequest) (*model.ProductSet, error) {
	resp := new(model.ProductSet)
	sb := sqlbuilder.NewBuilder("select * from p_product ")
	if req.OrderBy != "" {
		sb.OrderBy(req.OrderBy, req.Desc)
	}
	if req.Keywords != "" {
		sb.Where("name like ?", "%"+req.Keywords+"%")
	}

	if req.Size != 0 {
		//<offset>,<page_size>
		sb.Limit((req.Page-1)*req.Size, req.Size)
	}
	stmt, arg := sb.Build()

	var items = make([]*model.Product, 0, 0)
	err := p.db.Raw(stmt, arg...).Scan(&items).Error
	if err != nil {
		return nil, err
	}

	var count int64
	stmt2, arg2 := sb.BuildCount()
	err = p.db.Raw(stmt2, arg2...).Scan(&count).Error
	if err != nil {
		return nil, err
	}

	resp.Items = items
	resp.Total = count
	return resp, nil
}

func (p *ProductImpl) DescribeProduct(ctx context.Context, request *model.DescribeProductRequest) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductImpl) DeleteProduct(ctx context.Context, request *model.DeleteProductRequest) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProductImpl) UpdateProduct(ctx context.Context, request *model.UpdateProductRequest) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

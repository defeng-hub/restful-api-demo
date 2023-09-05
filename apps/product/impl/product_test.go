package impl

import (
	"context"
	"github.com/stretchr/testify/assert"
	"restful-api-demo/apps/product/model"
	"restful-api-demo/conf"
	"testing"
)

func init() {
	var err error
	//加载配置
	if err = conf.LoadConfigFromToml("../../../etc/pro.toml"); err != nil {
		panic("LoadConfigFromToml Fail")
	}

	if err = conf.LoadGlobalLogger(); err != nil {
		panic("LoadGlobalLogger Fail")
	}
}

// 表的注册
func TestDbAutoMigrate(t *testing.T) {
	should := assert.New(t)
	gdb, err := conf.C().MySQL.GetGormDB()
	should.NoError(err)

	err = gdb.AutoMigrate(&model.Product{})
	should.NoError(err)
	// 设置日志记录模式为详细模式
	gdb = gdb.Debug()
}

func TestProductImpl_SaveProduct(t *testing.T) {
	should := assert.New(t)
	db, err := conf.C().MySQL.GetGormDB()
	should.NoError(err)

	p := ProductImpl{
		db: db,
		l:  conf.L(),
	}
	product, err := p.SaveProduct(context.Background(), &model.Product{
		Name:  "产品3",
		Price: 300,
		Count: 3,
	})
	should.NoError(err)

	conf.L().Info(product)
}

func TestQueryProduct(t *testing.T) {
	//	QueryProduct
	should := assert.New(t)
	db, err := conf.C().MySQL.GetGormDB()
	should.NoError(err)
	db = db.Debug()
	p := ProductImpl{
		db: db,
		l:  conf.L(),
	}
	productset, err := p.QueryProduct(context.Background(), &model.QueryProductListRequest{
		OrderBy:  "name",
		Desc:     false,
		Page:     1,
		Size:     2,
		Keywords: "1",
	})
	should.NoError(err)

	conf.L().Info(productset)
}

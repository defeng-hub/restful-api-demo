package impl

import (
	"gorm.io/gorm"
	"restful-api-demo/apps/product"
	"restful-api-demo/common/logger"
	"restful-api-demo/conf"
)

type CouponImpl struct {
	db *gorm.DB
	l  logger.Logger
}

func (s *CouponImpl) Name() string {
	return product.AppName + "(" + couponService + ")"
}

func (s *CouponImpl) Config() {
	s.db, _ = conf.C().MySQL.GetGormDB()
	s.l = conf.L().Named(couponService)
}

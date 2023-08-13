package init

import (
	"restful-api-demo/apps"
	"restful-api-demo/apps/user"
	"restful-api-demo/conf"
)

type Init struct{}

var inits = new(Init)

func (s *Init) Config() {
	db, err := conf.C().MySQL.GetGormDB()
	if err != nil {
		panic("user模块 初始化失败")
	}
	RegisterTables(db)
}
func (s *Init) Name() string {
	return user.AppName
}

func init() {
	apps.RegistryImpl(inits)
}

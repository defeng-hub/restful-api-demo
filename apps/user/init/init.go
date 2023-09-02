package init

import (
	"restful-api-demo/apps"
	"restful-api-demo/apps/user"
	"restful-api-demo/conf"
)

type Init struct{}

func (s *Init) Config() {
	if db, err := conf.C().MySQL.GetGormDB(); err != nil {
		panic("user模块 初始化失败")
	} else {
		conf.L().Info(db)
		//RegisterTables(db)
	}

	InitRedis()
}
func (s *Init) Name() string {
	return user.AppName + "Init"
}

func init() {
	var inits = new(Init)
	apps.RegistryApp(inits)
}

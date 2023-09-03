package init

import (
	"github.com/go-redis/redis/v8"
	"restful-api-demo/apps"
	"restful-api-demo/apps/user"
	"restful-api-demo/conf"
)

type Init struct{}

var Rdb *redis.Client

func (s *Init) Config() {
	var err error
	if _, err := conf.C().MySQL.GetGormDB(); err != nil {
		panic("user模块 Mysql初始化失败")
	} else {
		//conf.L().Info(db)
		//RegisterTables(db)
	}

	if Rdb, err = InitRedis(); err != nil {
		panic("user模块 Redis初始化失败")
		return
	} else {
		conf.L().Infof("UserApp Redis:%s 加载成功...", Rdb)
	}

}
func (s *Init) Name() string {
	return user.AppName + "(Init)"
}

func init() {
	var inits = new(Init)
	apps.RegistryApp(inits)
}

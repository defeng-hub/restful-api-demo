package init

import (
	"github.com/go-redis/redis/v8"
	"restful-api-demo/apps"
	"restful-api-demo/apps/user"
	"restful-api-demo/apps/user/model"
	"restful-api-demo/common/cache/local_cache"
	"restful-api-demo/conf"
	"time"
)

type Init struct{}

var Rdb *redis.Client
var BlackCache local_cache.Cache

func (s *Init) Config() {
	var err error
	db, err := conf.C().MySQL.GetGormDB()
	if err != nil {
		panic("user模块 Mysql初始化失败")
	} else {
		//RegisterTables(db)
	}

	// 加载redis
	if Rdb, err = InitRedis(); err != nil {
		panic("user模块 Redis初始化失败")
		return
	} else {
		conf.L().Infof("UserApp Redis:%s 加载成功...", Rdb)
	}

	// 初始化黑名单Cache
	BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(time.Second * time.Duration(conf.C().Jwt.ExpiresTime)),
	)

	// 加载jwt黑名单
	var data []string
	err = db.Model(&model.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		conf.L().Error("加载数据库jwt黑名单失败!")
		return
	}
	for i := 0; i < len(data); i++ {
		BlackCache.SetDefault(data[i], struct{}{})
	}

}
func (s *Init) Name() string {
	return user.AppName + "(Init)"
}

func init() {
	var inits = new(Init)
	apps.RegistryApp(inits)
}

package init

import (
	"github.com/go-redis/redis/v8"
	"restful-api-demo/conf"
)

func InitRedis() (*redis.Client, error) {
	rdb, err := conf.C().Redis.GetRdb()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

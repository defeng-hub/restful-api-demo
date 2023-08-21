package init

import "restful-api-demo/conf"

func InitRedis() {
	_, err := conf.C().Redis.GetRdb()
	if err != nil {
		panic(err)
		return
	}
}

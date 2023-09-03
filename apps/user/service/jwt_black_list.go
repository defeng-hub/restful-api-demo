package service

import (
	"context"
	userdb "restful-api-demo/apps/user/init"
	"restful-api-demo/apps/user/model"
	"restful-api-demo/conf"
	"time"
)

type JwtService struct{}

//@description: 拉黑jwt
func (jwtService *JwtService) JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	db, err := conf.C().MySQL.GetGormDB()
	if err != nil {
		return err
	}
	err = db.Create(&jwtList).Error
	if err != nil {
		return err
	}

	// 本地缓存拉黑
	userdb.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@description: 判断JWT是否在黑名单内部
func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := userdb.BlackCache.Get(jwt)
	return ok
}

//@description: 从redis取jwt
func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = userdb.Rdb.Get(context.Background(), userName).Result()
	return err, redisJWT
}

//@description: jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(conf.C().Jwt.ExpiresTime) * time.Second
	//err = userdb.Rdb.Set(context.Background(), userName, jwt, timer).Err()

	err = userdb.Rdb.LPush(context.Background(), userName, jwt).Err()
	userdb.Rdb.Expire(context.Background(), userName, timer)
	return err
}

// 清除超出的token 并拉黑
func (jwtService *JwtService) RemoveExcessToken(userName string, num int64) error {
	lenn, err := userdb.Rdb.LLen(context.Background(), userName).Result()
	if err != nil || lenn <= num {
		return err
	}

	// 获取token， 拉黑token
	oldtoken, err := userdb.Rdb.LIndex(context.Background(), userName, num).Result()
	userdb.Rdb.LTrim(context.Background(), userName, 0, num-1)
	jwtService.JsonInBlacklist(model.JwtBlacklist{Jwt: oldtoken})
	return err
}

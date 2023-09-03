package service

import (
	"context"
	userdb "restful-api-demo/apps/user/init"
	"restful-api-demo/conf"
	"time"
)

type JwtService struct{}

//@description: 拉黑jwt
//func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
//	err = jwtService.db.Create(&jwtList).Error
//	if err != nil {
//		return
//	}
//	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
//	return
//}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool
//func (jwtService *JwtService) IsBlacklist(jwt string) bool {
//	_, ok := global.BlackCache.Get(jwt)
//	return ok
//	// err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
//	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
//	// return !isNotFound
//}

//@description: 从redis取jwt
func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = userdb.Rdb.Get(context.Background(), userName).Result()
	return err, redisJWT
}

//@description: jwt存入redis并设置过期时间
func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(conf.C().Jwt.ExpiresTime) * time.Second
	err = userdb.Rdb.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

//func LoadAll() {
//	var data []string
//	err := global.GVA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
//	if err != nil {
//		global.GVA_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
//		return
//	}
//	for i := 0; i < len(data); i++ {
//		global.BlackCache.SetDefault(data[i], struct{}{})
//	} // jwt黑名单 加入 BlackCache 中
//}

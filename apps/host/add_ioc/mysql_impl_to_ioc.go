package ioc

import (
	"restful-api-demo/apps"
	mysqlImpl "restful-api-demo/apps/host/impl"
)

// 这个对象要进入ioc模块,不过他当前并没有准备好, 需要对实例执行config方法
var impl = &mysqlImpl.MysqlServiceImpl{}

func init() {
	// 老方法都是在start的时候,手动把服务注册到IOC层,  案例: apps.HostService = impl.NewMysqlServiceImpl()
	// 现在采用自动注册,类似于mysql引擎   import _ "xxx"
	// sql这个库就是案例
	apps.RegistryImpl(impl)
}

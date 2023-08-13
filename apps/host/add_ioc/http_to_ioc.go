package ioc

import (
	"restful-api-demo/apps"
	"restful-api-demo/apps/host/http"
)

// 用来注册到ioc
var handler = &http.Handler{}

// 因为实现了上述的三个函数   所以可以注册进去ioc
func init() {
	apps.RegistryGin(handler)
}

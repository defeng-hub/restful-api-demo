package apps

import (
	"fmt"
	"github.com/defeng-hub/restful-api-demo/apps/host"
)

//IOC容器层: 管理所有的服务实例

// 1.HostService 的实例必须注册过来,注册完了才会有具体的实例
// 2.Http

var (
	HostService host.Service

	svcs = map[string]AppService{}
)

func Registry(obj AppService) {
	if _, ok := svcs[obj.Name()]; ok {
		panic(fmt.Sprintf("服务:%s 已经注册", obj.Name()))
	}
	// 服务实例注册到svcs map
	svcs[obj.Name()] = obj

	// 注册具体的服务
	if v, ok := obj.(host.Service); ok {
		HostService = v
	}
}
func AppsInit() {
	for _, svc := range svcs {
		svc.Config()
	}
}

type AppService interface {
	Config()
	Name() string
}

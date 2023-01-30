package apps

import "github.com/defeng-hub/restful-api-demo/apps/host"

//IOC容器层: 管理所有的服务实例

// 1.HostService 的实例必须注册过来,注册完了才会有具体的实例
// 2.Http

var (
	HostService host.Service
)

package apps

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//IOC容器层: 管理所有的服务实例
var (
	implApps = map[string]ImplService{}
	ginApps  = map[string]GinService{}
)

type ImplService interface {
	Config()
	Name() string
}

// GinService 注册由gin编写的handler
type GinService interface {
	Registry(r gin.IRouter)
	Name() string
	Config()
}

// RegistryImpl 注册实现类
func RegistryImpl(obj ImplService) {
	if _, ok := implApps[obj.Name()]; ok {
		panic(fmt.Sprintf("服务:%s 已经注册", obj.Name()))
	}
	// 服务实例注册到svcs map
	implApps[obj.Name()] = obj

}

// RegistryGin 注册到GinApps
func RegistryGin(obj GinService) {
	if _, ok := ginApps[obj.Name()]; ok {
		panic(fmt.Sprintf("服务:%s 已经注册", obj.Name()))
	}
	// 服务实例注册到svcs map
	ginApps[obj.Name()] = obj
}

// InitImpl 初始化全部Impl
func InitImpl() {
	for _, svc := range implApps {
		svc.Config()
	}
}

func InitGin(r gin.IRouter) {
	for _, v := range ginApps {
		// 先初始化对象
		v.Config()

		//	再完成http handler注册
		v.Registry(r)
	}

}

func GetImpl(name string) interface{} {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}
	return nil
}

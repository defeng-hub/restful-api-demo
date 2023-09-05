package all

import (
	//初始化
	_ "restful-api-demo/apps/user/init"

	// 案例模块
	_ "restful-api-demo/apps/host/add_ioc"

	// 用户模块
	_ "restful-api-demo/apps/user/add_ioc"

	// 系统模块
	_ "restful-api-demo/apps/system/add_ioc"

	// 产品模块
	_ "restful-api-demo/apps/product/add_ioc"
)

package all

import (
	_ "restful-api-demo/apps/user/init"

	_ "restful-api-demo/apps/host/add_ioc"

	//加载user实例
	_ "restful-api-demo/apps/user/add_ioc"

	//	 加载系统实例
	_ "restful-api-demo/apps/system/add_ioc"
)

package add_ioc

import (
	"restful-api-demo/apps"
	svr "restful-api-demo/apps/product/impl"
)

var ProductService = new(svr.ProductImpl)

func init() {
	apps.RegistryApp(ProductService)
}

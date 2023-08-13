package add_ioc

import (
	"restful-api-demo/apps"
	"restful-api-demo/apps/host/http"
)

var httpService = new(http.Handler)

func init() {
	apps.RegistryGin(httpService)
}

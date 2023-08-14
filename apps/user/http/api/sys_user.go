package api

import "restful-api-demo/apps/user/service"

type UserApi struct {
	Srv *service.UserService
}

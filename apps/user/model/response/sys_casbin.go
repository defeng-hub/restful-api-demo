package response

import (
	"restful-api-demo/apps/user/model/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths" comment:"casbin详情列表"`
}

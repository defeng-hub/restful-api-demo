package response

import "restful-api-demo/apps/user/model"

type SysAuthorityResponse struct {
	Authority model.SysAuthority `json:"authority" comment:"系统角色详情"`
}

type SysAuthorityCopyResponse struct {
	Authority      model.SysAuthority `json:"authority"  comment:"系统角色详情"`
	OldAuthorityId string             `json:"oldAuthorityId"  comment:"旧角色ID"` // 旧角色ID
}

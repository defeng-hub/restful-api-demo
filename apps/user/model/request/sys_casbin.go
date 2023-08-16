package request

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path" example:"/api"`  // 路径
	Method string `json:"method" example:"GET"` // 方法
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId string       `json:"authorityId"` // 权限id
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/UserService/GetUserList", Method: "GET"},
	}
}

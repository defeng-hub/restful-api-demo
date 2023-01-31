package host

import "context"

// Service  的接口定义
type Service interface {
	// CreateHost 录入主机
	SaveHost(context.Context, *Host) (*Host, error)
	// QueryHost 查询主机列表
	QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
}

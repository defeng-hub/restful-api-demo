package host

import "context"

type (
	// Service  的接口定义
	Service interface {
		// SaveHost 录入主机
		SaveHost(context.Context, *Host) (*Host, error)
		// QueryHost 查询主机列表
		QueryHost(context.Context, *QueryHostRequest) (*HostSet, error)
		// DescribeHost 查询主机详情
		DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)
		// UpdateHost 主机更新
		UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
		// DeleteHost 删除主机
		DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
	}
)

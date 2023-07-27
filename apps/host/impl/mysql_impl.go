package impl

import (
	"context"
	"database/sql"

	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/defeng-hub/restful-api-demo/conf"
	"github.com/defeng-hub/restful-api-demo/utils/sqlbuilder"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

type MysqlServiceImpl struct {
	db *sql.DB
	l  logger.Logger
}

//编译器做静态检查
var _ host.Service = (*MysqlServiceImpl)(nil)

// NewMysqlServiceImpl 创建mysql实现类的实例, 只给测试用例使用了
func NewMysqlServiceImpl() (*MysqlServiceImpl, error) {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return nil, err
	}
	return &MysqlServiceImpl{
		db: db,
		l:  zap.L().Named("Host Impl"),
	}, nil
}

func (s *MysqlServiceImpl) DescribeHost(ctx context.Context, request *host.QueryHostRequest) (*host.Host, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MysqlServiceImpl) UpdateHost(ctx context.Context, request *host.UpdateHostRequest) (*host.Host, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MysqlServiceImpl) DeleteHost(ctx context.Context, request *host.DeleteHostRequest) (*host.Host, error) {
	//TODO implement me
	panic("implement me")
}

func (s *MysqlServiceImpl) SaveHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	//校验参数
	if err := ins.Validate(); err != nil {
		return nil, err
	}
	s.l.Named("Create").Debug("create host")
	s.l.Infof("参数校验,%s", "通过")

	// dao层入库
	err := s.save(ctx, ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (s *MysqlServiceImpl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (
	*host.HostSet, error) {
	b := sqlbuilder.NewBuilder(queryHostSQL)
	if req.Keywords != "" {
		b.Where("r.`name`=? or r.`description`=?",
			"%"+req.Keywords+"%",
			"%"+req.Keywords+"%",
		)
	}
	b.Limit(req.OffSet(), uint(req.PageSize))
	querySQL, args := b.Build()
	//s.l.Infof("querysql:%s, args:%v", querySQL, args)

	stmt, err := s.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// id     | vendor | region      | create_at     | expire_at | type | name       |
	// description | status | update_at     | sync_at | account | public_ip |
	// private_ip | resource_id | cpu | memory | gpu_amount | gpu_spec | os_type | os_name | serial_number |
	set := host.NewHostSet()
	for rows.Next() {
		ins := host.NewHost()
		if err := rows.Scan(&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt, &ins.Type, &ins.Name,
			&ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt, &ins.Account, &ins.PublicIP, &ins.PrivateIP, &ins.Id,
			&ins.CPU, &ins.Memory, &ins.GPUAmount, &ins.GPUSpec, &ins.OSName, &ins.SerialNumber, &ins.SerialNumber); err != nil {
			return nil, err
		}
		set.Items = append(set.Items, ins)
	}
	for i := range set.Items {
		s.l.Infof("set:%v", set.Items[i].Name)
	}
	return set, nil
}

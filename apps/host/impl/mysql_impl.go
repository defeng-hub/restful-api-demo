package impl

import (
	"context"
	"database/sql"
	"fmt"

	"restful-api-demo/apps/host"
	"restful-api-demo/conf"
	"restful-api-demo/utils/sqlbuilder"

	"restful-api-demo/common/logger"
	"restful-api-demo/common/logger/zap"
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

func (i *MysqlServiceImpl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	b := sqlbuilder.NewBuilder(queryHostSQL)
	b.Where("r.id = ?", req.Id)

	querySQL, args := b.Build()
	i.l.Infof("describe sql: %s, args: %v", querySQL, args)

	// query stmt, 构建一个Prepare语句
	stmt, err := i.db.PrepareContext(ctx, querySQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	ins := host.NewHost()
	row := stmt.QueryRowContext(ctx, args...)
	err = row.Scan(
		&ins.Id, &ins.Vendor, &ins.Region, &ins.CreateAt, &ins.ExpireAt,
		&ins.Type, &ins.Name, &ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt,
		&ins.Account, &ins.PublicIP, &ins.PrivateIP, &ins.Id,
		&ins.CPU, &ins.Memory, &ins.GPUAmount, &ins.GPUSpec, &ins.OSType, &ins.OSName, &ins.SerialNumber,
	)
	if err != nil {
		return nil, err
	}

	return ins, nil
}

func (i *MysqlServiceImpl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	// 获取已有对象
	ins, err := i.DescribeHost(ctx, host.NewDescribeHostRequestWithId(req.Id))
	if err != nil {
		return nil, err
	}

	// 根据更新的模式, 更新对象
	switch req.UpdateMode {
	case host.UPDATE_MODE_PUT:
		if err := ins.Put(req.Host); err != nil {
			return nil, err
		}
		// 整个对象的局部更新
	case host.UPDATE_MODE_PATCH:
		if err := ins.Patch(req.Host); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("update_mode only requred put/patch")
	}

	// 检查更新后的数据是否合法
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	// 更新数据库里面的数据
	if err := i.update(ctx, ins); err != nil {
		return nil, err
	}

	// 返回更新后的对象
	return ins, nil
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
	b.Limit(req.OffSet(), int64(req.PageSize))
	querySQL, args := b.Build()
	s.l.Infof("生成的sql和参数:%s, args:%v", querySQL, args)
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

	// 构建count 查询
	countSQL, args := b.BuildCount()
	countStmt, _ := s.db.PrepareContext(ctx, countSQL)
	defer countStmt.Close()
	countStmt.QueryRowContext(ctx, args...).Scan(&set.Total)

	return set, nil
}

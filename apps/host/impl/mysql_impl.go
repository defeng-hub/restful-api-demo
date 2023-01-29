package impl

import (
	"context"
	"database/sql"
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/defeng-hub/restful-api-demo/conf"
)

//编译器做静态检查
var _ host.Service = (*MysqlServiceImpl)(nil)

type MysqlServiceImpl struct {
	DB *sql.DB
}

// NewMysqlServiceImpl 创建mysql实现类的实例,
func NewMysqlServiceImpl() (*MysqlServiceImpl, error) {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return nil, err
	}
	return &MysqlServiceImpl{
		DB: db,
	}, nil
}

func (s *MysqlServiceImpl) SaveHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	//校验参数
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	// dao层入库
	err := s.save(ctx, ins)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

func (s *MysqlServiceImpl) QueryHost(ctx context.Context, request *host.QueryHostRequest) (
	*host.HostSet, error) {
	return nil, nil
}

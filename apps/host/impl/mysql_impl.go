package impl

import (
	"context"
	"database/sql"
	"fmt"
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

func (s *MysqlServiceImpl) SaveHost(ctx context.Context, h *host.Host) (*host.Host, error) {
	fmt.Println("saveHost..............")
	fmt.Printf("#########\n参数1:%#v\n######\n", h)
	return nil, nil
}

func (s *MysqlServiceImpl) QueryHost(ctx context.Context, request *host.QueryHostRequest) (
	*host.HostSet, error) {
	fmt.Println("queryHost..............")
	return nil, nil
}

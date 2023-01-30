package impl

import (
	"context"
	"database/sql"
	"github.com/defeng-hub/restful-api-demo/apps"
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/defeng-hub/restful-api-demo/conf"
)

//编译器做静态检查
var _ host.Service = (*MysqlServiceImpl)(nil)

// 这个对象要进入ioc模块,不过他当前并没有准备好, 需要对实例执行config方法
var impl = &MysqlServiceImpl{}

type MysqlServiceImpl struct {
	DB *sql.DB
}

// NewMysqlServiceImpl 创建mysql实现类的实例, 只给测试用例使用了
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

// Config Name #####通过实现了下边两个方法就可以注册到ioc层了#####
func (i *MysqlServiceImpl) Config() {
	// 只需要保证config() 执行完成就能实现初始化
	i.DB, _ = conf.C().MySQL.GetDB()
}
func (i *MysqlServiceImpl) Name() string {
	return host.AppName
}

func init() {
	// 老方法都是在start的时候,手动把服务注册到IOC层,  案例: apps.HostService, _ = impl.NewMysqlServiceImpl()
	// 现在采用自动注册,类似于mysql引擎   import _ "xxx"
	// sql这个库就是案例
	//apps.HostService = impl // 注册到ioc层
	apps.Registry(impl)
}

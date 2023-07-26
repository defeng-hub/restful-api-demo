package impl_test

import (
	"context"
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/defeng-hub/restful-api-demo/apps/host/impl"
	"github.com/defeng-hub/restful-api-demo/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	service *impl.MysqlServiceImpl
)

func init() {
	conf.LoadConfigFromToml("../../../etc/pro.toml")
	//conf.LoadConfigFromEnv()
	conf.LoadGlobalLogger()
	service, _ = impl.NewMysqlServiceImpl()
}

func TestCreateHost(t *testing.T) {
	should := assert.New(t)
	host1 := host.NewHost()

	host1.Id = "111"
	host1.Region = "6"
	host1.Type = "6"
	host1.Name = "家用服务器"
	host1.CPU = 10
	host1.Memory = 20484

	host1, err := service.SaveHost(context.Background(), host1)
	if should.NoError(err) {
		should.NotNil(host1)
		println(host1)
	}
}

func TestQueryHost(t *testing.T) {
	_, err := service.QueryHost(context.Background(), host.NewQueryHostRequest(20, 1, ""))
	if err != nil {
		return
	}
}

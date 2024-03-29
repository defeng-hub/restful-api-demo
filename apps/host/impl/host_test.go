package impl_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"restful-api-demo/apps/host"
	"restful-api-demo/apps/host/impl"
	"restful-api-demo/conf"
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

	host1.Id = "444"
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
	should := assert.New(t)

	hostset, err := service.QueryHost(context.Background(), host.NewQueryHostRequest(6, 1, ""))
	should.NoError(err)
	should.NotNil(hostset)
	fmt.Println(hostset)
}

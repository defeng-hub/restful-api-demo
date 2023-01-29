package impl_test

import (
	"context"
	"fmt"
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/defeng-hub/restful-api-demo/apps/host/impl"
	"github.com/defeng-hub/restful-api-demo/conf"
	"testing"
)

var (
	service *impl.MysqlServiceImpl
)

func init() {
	conf.LoadConfigFromToml("../../../etc/pro.toml")
	service, _ = impl.NewMysqlServiceImpl()
}

func TestCreateHost(t *testing.T) {
	service.SaveHost(context.Background(), &host.Host{
		Resource: &host.Resource{
			Id:          "123",
			Vendor:      0,
			Region:      "",
			Type:        "",
			Name:        "",
			Description: "",
			Status:      "",
			Tags:        nil,
			UpdateAt:    0,
			SyncAt:      0,
			Account:     "",
			PublicIP:    "",
			PrivateIP:   "",
		},
		Describe: &host.Describe{
			CPU:          2,
			Memory:       1024,
			GPUAmount:    4,
			GPUSpec:      "4",
			OSType:       "linux",
			OSName:       "os",
			SerialNumber: "",
		},
	})
	err := service.DB.Ping()
	if err != nil {
		return
	}
	fmt.Println("pong.....")
}

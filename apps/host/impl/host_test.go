package impl_test

import (
	"context"
	"fmt"
	"github.com/defeng-hub/restful-api-demo/apps/host"
	"github.com/defeng-hub/restful-api-demo/apps/host/impl"
	"github.com/defeng-hub/restful-api-demo/conf"
	"testing"
	"time"
)

var (
	service *impl.MysqlServiceImpl
)

func init() {
	conf.LoadConfigFromToml("../../../etc/pro.toml")
	conf.LoadGlobalLogger()
	service, _ = impl.NewMysqlServiceImpl()
}

func TestCreateHost(t *testing.T) {
	_, err := service.SaveHost(context.Background(), &host.Host{
		Resource: &host.Resource{
			Id:          "1234",
			Vendor:      0,
			Region:      "6",
			Type:        "6",
			Name:        "家用服务器",
			Description: "def家",
			Status:      "6",
			UpdateAt:    time.Now().UnixMilli(),
			SyncAt:      0,
			Account:     "6",
			PublicIP:    "6",
			PrivateIP:   "6",
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
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = service.DB.Ping()
	if err != nil {
		return
	}
	fmt.Println("pong.....")
}

func TestQueryHost(t *testing.T) {
	_, err := service.QueryHost(context.Background(), host.NewQueryHostRequest(20, 1, ""))

	if err != nil {
		return
	}
}

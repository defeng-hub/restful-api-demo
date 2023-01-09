package conf_test

import (
	"fmt"
	"github.com/defeng-hub/restful-api-demo/conf"
	"os"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	conf.LoadConfigFromToml("../etc/demo.toml1")
	fmt.Printf("%#v", conf.C().App.Name)
}

func TestLoadconfigFromEnv(t *testing.T) {
	// 先设置 环境变量
	os.Setenv("MYSQL_USERNAME", "123")

	conf.LoadConfigFromEnv()
	fmt.Printf("%#v", conf.C().MySQL.UserName)
}

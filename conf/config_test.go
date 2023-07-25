package conf_test

import (
	"fmt"
	"github.com/defeng-hub/restful-api-demo/conf"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/pro.toml")
	if should.NoError(err) {
		should.Equal("demo1", conf.C().App.Name)
	}
}

func TestLoadconfigFromEnv(t *testing.T) {
	// 先设置 环境变量
	os.Setenv("MYSQL_USERNAME", "mcube-demo")
	should := assert.New(t)
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("mcube-demo", conf.C().MySQL.UserName)
	}
}

func TestGetDB(t *testing.T) {
	conf.LoadConfigFromToml("../etc/pro.toml")

	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		fmt.Printf("fail:%v", err)
		return
	}
	fmt.Printf("success:%#v", db)
}

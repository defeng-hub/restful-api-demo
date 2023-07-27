package conf_test

import (
	"github.com/defeng-hub/restful-api-demo/conf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/pro.toml")
	if should.NoError(err) {
		should.Equal("demo", conf.C().App.Name)
	}
}

func TestLoadconfigFromEnv(t *testing.T) {
	// 先设置 环境变量
	//oss.Setenv("MYSQL_USERNAME", "mcube-demo")
	should := assert.New(t)
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("mcube-demo", conf.C().MySQL.UserName)
	}
}

func TestGetDB(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/pro.toml")
	if should.NoError(err) {
		db, err := conf.C().MySQL.GetDB()
		should.NoError(err)
		should.NotNil(db)
	}

}

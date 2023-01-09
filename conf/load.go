package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

func LoadConfigFromToml(filePath string) error {
	config = NewDefaultConfig()
	_, err := toml.DecodeFile(filePath, config)
	if err != nil {
		return fmt.Errorf("打开配置文件出错:%s", err)
	}

	err2 := loadGloabl()
	if err2 != nil {
		return fmt.Errorf("连接数据库失败:%s", err)
	}
	return nil
}

func LoadConfigFromEnv() error {
	config = NewDefaultConfig()
	err := env.Parse(config)
	if err != nil {
		return fmt.Errorf("获取env环境变量出错:%s", err)
	}
	err2 := loadGloabl()
	if err2 != nil {
		return fmt.Errorf("连接数据库失败:%s", err)
	}
	return nil
}

// 全局实例
func loadGloabl() (err error) {
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return
	}
	return
}

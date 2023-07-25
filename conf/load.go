package conf

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

// LoadConfigFromToml 通过toml获取配置
func LoadConfigFromToml(filePath string) error {
	config = newDefaultConfig()
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

// LoadConfigFromEnv 通过env获取配置
func LoadConfigFromEnv() error {
	config = newDefaultConfig()
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

// 加载全局实例, 给上边两个用
func loadGloabl() (err error) {
	db, err = config.MySQL.getDBConn()
	if err != nil {
		return
	}
	return
}

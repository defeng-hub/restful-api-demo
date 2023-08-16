package conf

import (
	"fmt"
)

func (a *App) HttpAddr() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

func newDefaultConfig() *Config {
	return &Config{
		App:   newDefaultApp(),
		MySQL: newDefaultMysql(),
		Log:   newDefaultLog(),
	}
}

func newDefaultApp() *App {
	return &App{
		Name: "demoapp",
		Host: "127.0.0.1",
		Port: "8050",
		Key:  "key",
	}
}

func newDefaultMysql() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "test",
		Password:    "password",
		Database:    "db1",
		MaxOpenConn: 50,
		MaxIdleConn: 20,
		MaxLifeTime: 1800,
		MaxIdleTime: 600,
	}
}

// newDefaultLog
func newDefaultLog() *Log {
	return &Log{
		Level:  "debug",
		Format: "text",
		To:     "stdout",
	}
}

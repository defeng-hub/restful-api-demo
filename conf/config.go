package conf

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //mysql 驱动
	"sync"
	"time"
)

// 为了不让程序在运行时恶意修改,设置成私有变量
var (
	config *Config
	db     *sql.DB
)

// C 想要从外部获取配置, 通过C获取config对象
func C() *Config {
	return config
}

// Config 应用配置
type Config struct {
	App   *App   `toml:"app"`
	MySQL *MySQL `toml:"mysql"`
}

func newDefaultConfig() *Config {
	return &Config{
		App:   newDefaultApp(),
		MySQL: newDefaultMysql(),
	}
}

func newDefaultApp() *App {
	return &App{
		Name: "app",
		Host: "127.0.0.1",
		Port: "8099",
		Key:  "Key",
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
func (a *App) HttpAddr() string {
	return fmt.Sprintf("%s:%s", a.Host, a.Port)
}

type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
	Key  string `toml:"key" env:"APP_KEY"`
}

type MySQL struct {
	Host     string `toml:"host" env:"MYSQL_HOST"`
	Port     string `toml:"port" env:"MYSQL_PORT"`
	UserName string `toml:"username" env:"MYSQL_USERNAME"`
	Password string `toml:"password" env:"MYSQL_PASSWORD"`
	Database string `toml:"database" env:"MYSQL_DATABASE"`
	// mySQL当前程序的连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`

	// 控制MySQL复用, 比如5,最多运行5个来赋予
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	// 一个连接的生命周期, 比如设计1h, 1h后换一个conn,保证可用
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	// 一个连接最长存活时间
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`
	lock        sync.Mutex
}

func (m *MySQL) GetDB() (*sql.DB, error) {
	// 直接锁住临界区,
	m.lock.Lock()
	defer m.lock.Unlock()
	// 加载全局数据库单例
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}

// getDBConn 获取数据库连接池
func (m *MySQL) getDBConn() (*sql.DB, error) {
	var err error
	//multiStatements=true 运行执行多行sql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true", m.UserName, m.Password, m.Host, m.Port, m.Database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}

package conf

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"restful-api-demo/common/logger/zap"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql" //mysql 驱动
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 为了不让程序在运行时恶意修改,设置成私有变量
var (
	config *Config
	db     *sql.DB
	gdb    *gorm.DB
	rdb    *redis.Client
)

// C 想要从外部获取配置, 通过C获取config对象
func C() *Config {
	return config
}
func L() *zap.Logger {
	return zap.L()
}

// Config 应用配置
type Config struct {
	App   *App   `toml:"app"`
	MySQL *MySQL `toml:"mysql"`
	Redis *Redis `toml:"redis"`
	Log   *Log   `toml:"log"`
	Jwt   *Jwt   `toml:"jwt"`
}

type Log struct {
	Level  string    `toml:"level" env:"LOG_LEVEL"`
	OutDir string    `toml:"out_dir" env:"LOG_PATH_DIR"`
	Format LogFormat `toml:"format" env:"LOG_FORMAT"`
	To     LogTo     `toml:"to" env:"LOG_TO"`
}
type App struct {
	Name string `toml:"name" env:"APP_NAME"`
	Host string `toml:"host" env:"APP_HOST"`
	Port string `toml:"port" env:"APP_PORT"`
	Key  string `toml:"key" env:"APP_KEY"`
}
type Jwt struct {
	SigningKey  string `toml:"signing_key"`  // jwt签名
	ExpiresTime int64  `toml:"expires_time"` // 过期时间
	BufferTime  int64  `toml:"buffer_time"`  // 缓冲时间
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
	// 1.第一种方式，使用LoadGlobal去加载 [伪代码 db = getDBConn()]， 在程序启动时，初始化全局db实例
	// 2.第二种方式，惰性加载
	m.lock.Lock() // 直接锁住临界区,
	defer m.lock.Unlock()
	// 加载全局数据库单例,以下内容加了锁，只能保证同时只有一个进来
	if db == nil {
		conn, err := m.getDBConn()
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}

func (m *MySQL) GetGormDB() (*gorm.DB, error) {
	if gdb == nil {
		// 加载gorm db
		sqlDB, err := m.GetDB()
		if err != nil {
			return nil, err
		}
		gdb, err = gorm.Open(mysql.New(mysql.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		return gdb, nil
	}
	return gdb, nil
}

// 获取数据库连接池  对内!!
func (m *MySQL) getDBConn() (*sql.DB, error) {
	var err error
	//multiStatements=true 运行执行多行sql
	//一个数据库链接配置：charset=utf8mb4&parseTime=True&loc=Local
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true&charset=utf8mb4&parseTime=True&loc=Local", m.UserName, m.Password, m.Host, m.Port, m.Database)
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

type Redis struct {
	DB       int    `toml:"db" env:"REDIS_DB"`       // redis的哪个数据库
	Addr     string `toml:"addr" env:"REDIS_ADDR"`   // 服务器地址:端口
	Password string `toml:"password" env:"REDIS_PW"` // 密码
	lock     sync.Mutex
}

func (r *Redis) initRedis() error {
	client := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password, // no password set
		DB:       r.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		rdb = nil
		return fmt.Errorf("redis load fail:, err:%v", err)
	} else {
		rdb = client
	}
	return nil

}

func (r *Redis) GetRdb() (*redis.Client, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	if rdb == nil {
		err := r.initRedis()
		if err != nil {
			L().Named("Init").Errorf("redis load fail: %v", err)
			return nil, err
		}
		return rdb, nil
	} else {
		return rdb, nil
	}
}

package conf

// 为了不让程序在运行时恶意修改,设置成私有变量
var config *Config

// C 想要获取配置, 通过C获取config对象
func C() *Config {
	return config
}

// Config 应用配置
type Config struct {
	App   *App   `toml:"app"`
	MySQL *MySQL `toml:"mysql"`
}

func NewDefaultConfig() *Config {
	return &Config{
		App:   NewDefaultApp(),
		MySQL: NewDefaultMysql(),
	}
}

func NewDefaultApp() *App {
	return &App{
		Name: "app",
		Host: "127.0.0.1",
		Port: "8099",
		Key:  "Key",
	}
}

func NewDefaultMysql() *MySQL {
	return &MySQL{
		Host:        "127.0.0.1",
		Port:        "3306",
		UserName:    "test",
		Password:    "password",
		Database:    "db1",
		MaxOpenConn: 50,
		MaxIdleConn: 30,
		MaxLifeTime: 3600,
		MaxIdleTime: 3600 * 2,
	}
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
	// mySQL连接数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	// 控制MySQL复用
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_idle_TIME"`
}

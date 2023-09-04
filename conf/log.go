package conf

import (
	"fmt"
	"restful-api-demo/common/logger/zap"
	"time"
)

// LogFormat 日志格式
type LogFormat string

const (
	// TextFormat 文本格式
	TextFormat = LogFormat("text")
	// JSONFormat json格式
	JSONFormat = LogFormat("json")
)

// LogTo 日志记录到哪儿
type LogTo string

const (
	// ToFile 保存到文件
	ToFile = LogTo("file")
	// ToStdout 打印到标准输出
	ToStdout = LogTo("stdout")
)

// LoadGlobalLogger  为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func LoadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)
	//根据config 获取全局logger对象
	lc := C().Log

	//设置日志级别
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("未获取到日志级别,将使用的Info日志级别:%s", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("Logger级别: %s", lv)
	}

	// 使用默认配置, 设置日志级别
	zapConfig := zap.DefaultConfig()
	zapConfig.Level = level                // 配置日志的level级别
	zapConfig.Files.RotateOnStartup = true // 程序每启动一次, 生成一个新的日志文件

	// 配置日志的输出方式
	switch lc.To {
	case ToStdout: //标准输出
		zapConfig.ToStderr = true // 输出到控制台
		zapConfig.ToFiles = false // 不输出到文件
	case ToFile:
		zapConfig.ToStderr = false
		zapConfig.ToFiles = true
		time_s := time.Now().Format("2006-01-02-15:04:05")
		//输出到文件
		zapConfig.Files.Name = time_s + ".log"
		zapConfig.Files.Path = lc.OutDir
		zapConfig.Files.MaxSize = 10 * 1024 * 1024
		zapConfig.Files.MaxBackups = 7

	default:
		zapConfig.ToStderr = true // 输出到控制台
		zapConfig.ToFiles = false // 不输出到文件
	}
	// 配置日志输出格式
	switch lc.Format {
	case TextFormat:
		break
	case JSONFormat:
		zapConfig.JSON = true
	default: //默认就是 TextFormat 'text'
		break
	}
	// 把配置应用到全局
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}

	//输出log初始化 是否完成的日志
	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

package cmd

import (
	"github.com/defeng-hub/restful-api-demo/apps"
	"github.com/defeng-hub/restful-api-demo/conf"
	"github.com/defeng-hub/restful-api-demo/protocol"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"

	//注册所有的服务实例
	_ "github.com/defeng-hub/restful-api-demo/all"
)

var (
	confFile string
)

// StartCmd 程序的启动时, 组装在这里进行
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动服务",
	Long:  `将启动grpc和http对外进行服务`,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		//加载配置
		if err = conf.LoadConfigFromToml(confFile); err != nil {
			return err
		}

		if err = conf.LoadGlobalLogger(); err != nil {
			return err
		}

		//对所有app-impl进行初始化
		apps.InitImpl()

		//完成暴露http服务的初始化
		//g := gin.Default()
		//apps.InitGin(g)
		//g.Run(conf.C().App.HttpAddr())
		manage := newManage()
		//启动http服务
		if err = manage.http.Start(); err != nil {
			manage.l.Errorf("http服务启动失败%s", err)
			return err
		}

		ch := make(chan os.Signal, 1)
		// channel是一种复合数据结构, 可以当初一个容器, 自定义的struct make(chan int, 1000), 8bytes * 1024  1Kb
		// 如果没close gc是不会回收的
		defer close(ch)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP, syscall.SIGINT)
		go manage.WaitStop(ch)
		return err
	},
}

func newManage() *manager {
	return &manager{
		http: protocol.NewHttpService(),
		l:    zap.L().Named("Global Manager"),
	}
}

// 处理来自外部的中断信号, 比如Terminal
func (m *manager) WaitStop(ch <-chan os.Signal) {
	for v := range ch {
		switch v {
		default:
			m.l.Infof("received signal: %s", v)
			m.http.Stop()
		}
	}
}

// 管理http
type manager struct {
	http *protocol.HTTPService
	l    logger.Logger
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f",
		"etc/pro.toml", "配置文件路径")
	RootCmd.AddCommand(StartCmd)
}

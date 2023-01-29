package cmd

import (
	"github.com/defeng-hub/restful-api-demo/apps/host/http"
	"github.com/defeng-hub/restful-api-demo/apps/host/impl"
	"github.com/defeng-hub/restful-api-demo/conf"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	// pusher service config option
	confType string
	confFile string
	confETCD string
)

// StartCmd 程序的启动时, 组装在这里进行
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 demo-api 管理系统",
	Long:  `启动 demo-api ...`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//加载配置
		if err := conf.LoadConfigFromToml(confFile); err != nil {
			return err
		}
		service, _ := impl.NewMysqlServiceImpl()

		api := http.NewHttpHandler(service)

		g := gin.Default()
		api.Registry(g)

		return g.Run(conf.C().App.HttpAddr())
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f",
		"etc/pro.toml", "配置文件路径")
	RootCmd.AddCommand(StartCmd)
}

package cmd

import (
	"github.com/defeng-hub/restful-api-demo/apps"
	"github.com/defeng-hub/restful-api-demo/conf"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

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
	RunE: func(cmd *cobra.Command, args []string) error {
		//加载配置
		if err := conf.LoadConfigFromToml(confFile); err != nil {
			return err
		}

		if err := conf.LoadGlobalLogger(); err != nil {
			return err
		}

		//对所有app-impl进行初始化
		apps.InitImpl()

		//完成暴露http服务的初始化
		g := gin.Default()
		apps.InitGin(g)

		return g.Run(conf.C().App.HttpAddr())
	},
}

func init() {
	StartCmd.PersistentFlags().StringVarP(&confFile, "config", "f",
		"etc/pro.toml", "配置文件路径")
	RootCmd.AddCommand(StartCmd)
}

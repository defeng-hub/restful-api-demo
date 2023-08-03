package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"restful-api-demo/version"
)

var (
	vers bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "restful-api-demo.exe",
	Short: "主文件(restful-api-demo)",
	Long:  `这个程序是王丰的练习程序...`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
		}

		return nil
	},
}

func init() {

	//选项分为两种

	// 第一种永久选项 global flag
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "输出版本信息")
	RootCmd.AddCommand(StartCmd)

}

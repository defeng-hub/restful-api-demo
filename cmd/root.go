package cmd

import (
	"fmt"
	"github.com/defeng-hub/restful-api-demo/version"
	"github.com/spf13/cobra"
)

var (
	vers bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo短描述",
	Long:  `这是demo的长描述,可以很长.....`,
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
}

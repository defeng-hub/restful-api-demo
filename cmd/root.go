package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	version bool
	info    string
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "demo",
	Short: "demo短描述",
	Long:  `这是demo的长描述,可以很长.....`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("demo v1.0.0")
		}
		if info != "" {
			fmt.Println(info)
		}
		return nil
	},
}

func init() {

	//选项分为两种

	// 第一种永久选项 global flag
	RootCmd.PersistentFlags().BoolVarP(&version, "version", "v", false, "输出版本信息")
	// 第二种 本地选项，只能在定义它的命令中使用
	RootCmd.Flags().StringVarP(&info, "info", "i", "", "原样输出您的输入")
}

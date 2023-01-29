package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "demo-api",
	Short: "demo-api 管理系统",
	Long:  `demo-api ...`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("no flags find -root")
	},
}

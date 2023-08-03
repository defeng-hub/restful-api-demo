package main

import (
	"fmt"
	"restful-api-demo/cmd"
)

func main() {
	//加载配置
	//conf.LoadConfigFromToml("./etc/dev.toml")
	//fmt.Printf("%#v", conf.C().MySQL.Password)
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}

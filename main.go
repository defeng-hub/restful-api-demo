package main

import (
	"fmt"
	"github.com/defeng-hub/restful-api-demo/cmd"
)

func main() {
	//加载配置
	//conf.LoadConfigFromToml("./etc/dev.toml")
	//fmt.Printf("%#v", conf.C().MySQL.Password)
	// wdf分支 加一句话，第一次提交
	// wdf分支 加一句话，第二次提交
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

}

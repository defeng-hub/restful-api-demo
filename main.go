package main

import (
	"fmt"
	"github.com/defeng-hub/restful-api-demo/conf"
)

func main() {
	//加载配置
	conf.LoadConfigFromToml("./etc/dev.toml")
	fmt.Printf("%#v", conf.C().MySQL.Password)

}

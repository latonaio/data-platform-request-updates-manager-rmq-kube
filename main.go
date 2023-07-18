package main

import (
	"data-platform-request-updates-manager-rmq-kube/config"
	_ "data-platform-request-updates-manager-rmq-kube/routers"

	"github.com/astaxie/beego"
)

func main() {
	conf := config.NewConf()
	beego.Run(conf.SERVER.ServerURL())
}

package main

import (
	"OnlineResume-GoVersion/init_conf"
	_ "OnlineResume-GoVersion/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}

	// 配置初始化
	init_conf.Init()

	beego.Run()
}

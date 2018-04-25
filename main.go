package main

import (
	_ "usercenter/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		//beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		//logs.SetLogger(logs.AdapterConn, `{"net":"tcp", "addr":"0.0.0.0:7020"}`)
		beego.BeeLogger.SetLevel(beego.LevelDebug)
	}
	// init log
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"/tmp/uc/logs/uc.log", "maxdays":30}`)
	beego.Run()

}

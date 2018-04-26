package main

import (
	_ "usercenter/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	orm.DefaultTimeLoc = time.UTC
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error(err)
		panic(err)
	}
	dbUrl := beego.AppConfig.String("DBUrl")
	err = orm.RegisterDataBase("default", "mysql", dbUrl)
	if err != nil {
		beego.Error(err)
		panic(err)
	}
}

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

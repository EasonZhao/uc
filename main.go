package main

import (
	_ "usercenter/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func init() {
	// init log
	os.Mkdir("/tmp/uc", 0777)
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"/tmp/uc/debug.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":30}`)

	// init orm
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbUrl := beego.AppConfig.String("DBUrl")
	err := orm.RegisterDataBase("default", "mysql", dbUrl)
	if err != nil {
		logs.Critical(err)
	}
}

func main() {
	orm.RunCommand()
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BeeLogger.SetLevel(beego.LevelDebug)
	}
	beego.Run()
}

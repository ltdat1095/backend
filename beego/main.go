package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ltdat/backend/routers"
)

func init() {
	mysqluser := beego.AppConfig.String("mysqluser")
	mysqlpass := beego.AppConfig.String("mysqlpass")
	mysqlhost := beego.AppConfig.String("mysqlhost")
	mysqlport, _ := beego.AppConfig.Int("mysqlport")
	mysqldb := beego.AppConfig.String("mysqldb")

	connectString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		mysqluser, mysqlpass, mysqlhost, mysqlport, mysqldb)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", connectString)

	// dbName := "default"
	// force := false
	// verbose := false
	// err := orm.RunSyncdb(dbName, force, verbose)
	// if err != nil {
	// 	beego.Alert("Error in SyncDB", err)
	// }

	beego.SetStaticPath("/static", "static")
}

func main() {
	beego.Run()
}

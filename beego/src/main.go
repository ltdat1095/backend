package main

import (
	_ "backend/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		beego.AppConfig.String("mysqluser"),
		beego.AppConfig.String("mysqlpass"),
		beego.AppConfig.String("mysqlurls"),
		beego.AppConfig.String("mysqldb"))

	driverName := "mysql"
	dbName := "default"
	force := false
	verbose := false

	orm.RegisterDriver(driverName, orm.DRMySQL)
	orm.RegisterDataBase("default", driverName, connectString)

	err := orm.RunSyncdb(dbName, force, verbose)
	if err != nil {
		beego.Alert("Error in SyncDB", err)
	}
}

func main() {
	beego.Run()
}

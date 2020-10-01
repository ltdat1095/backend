package main

import (
	"fmt"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ltdat/backend/routers"
)

func init() {
	// Beego config
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.SetStaticPath("/static", "static")

	// Connect database
	mysqlhost := os.Getenv("RDS_HOSTNAME")
	mysqlport := os.Getenv("RDS_PORT")
	mysqluser := os.Getenv("RDS_USERNAME")
	mysqlpass := os.Getenv("RDS_PASSWORD")
	mysqldb := os.Getenv("RDS_DB_NAME")

	connectFormat := "%s:%s@tcp(%s:%s)/%s?charset=utf8"
	connectString := fmt.Sprintf(connectFormat,
		mysqluser, mysqlpass, mysqlhost, mysqlport, mysqldb)
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", connectString)

	dbName := "default"
	force := false
	verbose := false
	err := orm.RunSyncdb(dbName, force, verbose)
	if err != nil {
		beego.Alert("Error in SyncDB", err)
	}
}

func main() {
	sessionconf := &session.ManagerConfig{
		CookieName: "beegoSession",
		Gclifetime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory", sessionconf)
	go beego.GlobalSessions.GC()

	beego.Run()
}

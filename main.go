package main

import (
	"clserp/clserp/auth"
	_ "hello/routers"

	"github.com/astaxie/beego"
)

func main() {
	rltcfg := auth.MySQLConfig{
		Address: beego.AppConfig.String("mysqlhost"),
		Port:    beego.AppConfig.String("mysqlport"),
		DB:      beego.AppConfig.String("mysqldatabase"),
		User:    beego.AppConfig.String("mysqluser"),
		Pwd:     beego.AppConfig.String("mysqlpassword"),
	}
	auth.Init(rltcfg)
	beego.Run()
}

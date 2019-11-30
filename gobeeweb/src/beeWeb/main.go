package main

import (
	_ "beeWeb/routers"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	// beego默认配置
	//beego.BConfig.AppName = "beegoproject"
	beego.BConfig.RouterCaseSensitive = true //true，区分大小写

	beego.BConfig.Log.AccessLogs = true //是否输出日志到 Log
	beego.BConfig.Log.Outputs["console"] = ""

	beego.SetStaticPath("/downloads", "downloads")
	// 项目配置
	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPass := beego.AppConfig.String("mysqlpass")
	mysqlUrls := beego.AppConfig.String("mysqlurls")
	mysqlDb := beego.AppConfig.String("mysqldb")
	fmt.Printf("user=%s,pass=%s,url=%s,db=%s\n", mysqlUser, mysqlPass, mysqlUrls, mysqlDb)
	beego.Run()
}

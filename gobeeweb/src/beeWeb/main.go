package main

import (
	_ "beeWeb/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/howeyc/fsnotify"
	"log"
)

func main() {
	//beego默认配置
	//beego.BConfig.AppName = "beegoproject"
	beego.BConfig.RouterCaseSensitive = true //true，区分大小写

	beego.BConfig.Log.AccessLogs = true //是否输出日志到 Log
	beego.BConfig.Log.Outputs["console"] = ""

	// 可以用用静态文件配置
	beego.SetStaticPath("/downloads", "downloads")
	// 项目配置
	mysqlUser := beego.AppConfig.String("mysqluser")
	mysqlPass := beego.AppConfig.String("mysqlpass")
	mysqlUrls := beego.AppConfig.String("mysqlurls")
	mysqlDb := beego.AppConfig.String("mysqldb")
	fmt.Printf("user=%s,pass=%s,url=%s,db=%s\n", mysqlUser, mysqlPass, mysqlUrls, mysqlDb)
	beego.Run()
}






func logWeb()  {

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Watch("testDir")
	if err != nil {
		log.Fatal(err)
	}

	// Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	watcher.Close()
}

package main

import (
	"goblog/core"
	"goblog/flags"
	"goblog/global"
	"goblog/router"
	logservice "goblog/service/log_service"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	flags.Run()

	log := logservice.NewRuntimeLog("同步文章数据", logservice.RuntimeDataHour)
	log.SetItem("文章 1", 11)
	log.Save()
	log.SetItem("文章 2", 22)
	log.Save()
	//启动 web 程序
	router.Run()
}

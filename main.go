package main

import (
	"goblog/core"
	"goblog/flags"
	"goblog/global"
	"goblog/router"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	flags.Run()

	//启动 web 程序
	router.Run()
}

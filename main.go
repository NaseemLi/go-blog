package main

import (
	"goblog/core"
	"goblog/flags"
	"goblog/global"
	"goblog/router"
	cronservice "goblog/service/cron_service"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	global.Redis = core.InitRedis()
	global.ESClient = core.EsConnect()

	flags.Run()

	core.InitMysqlSE()

	//定时任务
	cronservice.Cron()

	//启动 web 程序
	router.Run()
}

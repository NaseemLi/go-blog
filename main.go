package main

import (
	"goblog/core"
	"goblog/flags"
	"goblog/global"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.ReadConf()
	core.InitLogrus()
	global.DB = core.InitDB()
	flags.Run()
}

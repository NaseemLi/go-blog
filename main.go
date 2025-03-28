package main

import (
	"fmt"
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
	fmt.Println("数据库连接成功")
}

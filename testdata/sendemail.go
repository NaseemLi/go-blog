package main

import (
	"fmt"
	"goblog/core"
	"goblog/flags"
	"goblog/global"
	emailservice "goblog/service/email_service"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()

	emailservice.SendRegisterCode("l635220481@126.com", "1234")
	fmt.Println("发送成功")
}

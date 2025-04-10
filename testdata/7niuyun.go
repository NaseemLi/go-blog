package main

import (
	"fmt"
	"goblog/core"
	"goblog/flags"
	"goblog/global"
	qiniuservice "goblog/service/qiniu_service"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()
	// url, err := SendFile("uploads/images/WechatIMG33.jpg")
	// fmt.Println(url, err)

	// file, _ := os.Open("uploads/images/WechatIMG33.jpg")
	// url, err := SendReader(file)
	// fmt.Println(url, err)

	fmt.Println(qiniuservice.GenToken())
}

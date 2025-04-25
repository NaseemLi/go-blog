package core

import (
	"goblog/global"
	river "goblog/service/river_service"

	"github.com/sirupsen/logrus"
)

func InitMysqlSE() {
	if !global.Config.River.Enable {
		logrus.Infof("关闭 mysql 同步操作")
		return
	}
	if !global.Config.ES.Enable {
		logrus.Infof("未配置es,关闭 mysql 同步操作")
		return
	}

	r, err := river.NewRiver()
	if err != nil {
		logrus.Fatal(err)
	}
	go r.Run()
}

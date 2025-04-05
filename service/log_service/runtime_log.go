package logservice

import (
	"goblog/global"
	"goblog/models"
	"goblog/models/enum"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RuntimeLog struct {
	level           enum.LogLevelType
	title           string
	itemList        []string
	serviceName     string
	runtimeDataType RuntimeDataType
}

func (r RuntimeLog) Save() {
	//判断创建还是更新
	var log models.LogModel

	var query gorm.DB
	switch r.runtimeDataType {
	case RuntimeDataHour:
		query.Where("")
	case RrntimeDataDay:
		query.Where("date (create_at) = date(now())")
	case RrntimeDataWeek:
	case RrntimeDataMouth:
	}

	global.DB.Where(query).Find(&log, "service_name = ? and log_type = ?", r.serviceName, enum.RuntimeLogType)

	content := strings.Join(r.itemList, "\n")

	if log.ID != 0 {
		//更新
		return
	}

	err := global.DB.Create(&models.LogModel{
		LogType: enum.RuntimeLogType,
		Title:   r.title,
		Content: content,
		Level:   r.level,
	}).Error
	if err != nil {
		logrus.Errorf("创建运行日志错误失败 %s", err)
		return
	}
}

type RuntimeDataType int8

const (
	RuntimeDataHour  RuntimeDataType = 1
	RrntimeDataDay   RuntimeDataType = 2 //按天分割
	RrntimeDataWeek  RuntimeDataType = 3
	RrntimeDataMouth RuntimeDataType = 4
)

func NewRuntimeLog(serviceName string, dateType RuntimeDataType) *RuntimeLog {
	return &RuntimeLog{
		serviceName:     serviceName,
		runtimeDataType: dateType,
	}
}

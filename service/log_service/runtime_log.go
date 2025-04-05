package logservice

import (
	"encoding/json"
	"fmt"
	"goblog/global"
	"goblog/models"
	"goblog/models/enum"
	"reflect"
	"strings"

	e "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type RuntimeLog struct {
	level           enum.LogLevelType
	title           string
	itemList        []string
	serviceName     string
	runtimeDataType RuntimeDataType
}

func (ac RuntimeLog) Save() {
	//判断创建还是更新
	var log models.LogModel

	global.DB.Find(&log, fmt.Sprintf(
		"service_name = ? and log_type = ? and created_at >= DATE_SUB(NOW(), %s)",
		ac.runtimeDataType.GetSqlTime()), ac.serviceName, enum.RuntimeLogType)
	content := strings.Join(ac.itemList, "\n")

	if log.ID != 0 {
		//更新
		c := strings.Join(ac.itemList, "\n")
		newContent := log.Content + "\n" + c

		//之前已经 save 过,那就更新
		global.DB.Model(&log).Updates(map[string]any{
			"content": newContent,
		})
		ac.itemList = []string{}
		return
	}

	err := global.DB.Create(&models.LogModel{
		LogType:     enum.RuntimeLogType,
		Title:       ac.title,
		Content:     content,
		Level:       ac.level,
		ServiceName: ac.serviceName,
	}).Error
	if err != nil {
		logrus.Errorf("创建运行日志错误失败 %s", err)
		return
		ac.itemList = []string{}
	}
}

func (ac *RuntimeLog) SetTitle(title string) {
	ac.title = title
}

func (ac *RuntimeLog) SetLevel(level enum.LogLevelType) {
	ac.level = level
}

func (ac *RuntimeLog) SetLink(label string, href string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_item link\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\"><a href=\"%s\" target=\"_blank\">%s</a></div></div>",
		label,
		href, href))
}

func (ac *RuntimeLog) SetImage(src string) {
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_image\"><img src=\"%s\" alt=\"\"></div>", src))
}

func (ac *RuntimeLog) setItem(label string, value any, LogLevelType enum.LogLevelType) {
	var v string

	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Struct, reflect.Map, reflect.Slice:
		byteData, _ := json.Marshal(value)
		v = string(byteData)
	default:
		v = fmt.Sprintf("%v", value)
	}

	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_item %s\"><div class=\"log_item_label\">%s</div><div class=\"log_item_content\">%s</div></div>",
		LogLevelType,
		label, v))
}

func (ac *RuntimeLog) SetItem(label string, value any) {
	ac.setItem(label, value, enum.LogInfoLevel)
}

func (ac *RuntimeLog) SetItemInfo(label string, value any) {
	ac.setItem(label, value, enum.LogInfoLevel)
}

func (ac *RuntimeLog) SetItemWarn(label string, value any) {
	ac.setItem(label, value, enum.LogWarnLevel)
}

func (ac *RuntimeLog) SetItemError(label string, value any) {
	ac.setItem(label, value, enum.LogErrLevel)
}

func (ac *RuntimeLog) SetError(label string, err error) {
	msg := e.WithStack(err)
	logrus.Errorf("%s %s", label, err.Error())
	ac.itemList = append(ac.itemList, fmt.Sprintf("<div class=\"log_error\"><div class=\"line\"><div class=\"label\">%s</div><div class=\"value\">%s</div><div class=\"type\">%T</div></div><div class=\"stack\">%+v</div></div>",
		label,
		err,
		err,
		msg))
}

type RuntimeDataType int8

const (
	RuntimeDataHour  RuntimeDataType = 1
	RrntimeDataDay   RuntimeDataType = 2 //按天分割
	RrntimeDataWeek  RuntimeDataType = 3
	RrntimeDataMouth RuntimeDataType = 4
)

func (r RuntimeDataType) GetSqlTime() string {
	switch r {
	case RuntimeDataHour:
		return "interval 1 HOUR"
	case RrntimeDataDay:
		return "interval 1 DAY"
	case RrntimeDataWeek:
		return "interval 1 WEEK"
	case RrntimeDataMouth:
		return "interval 1 MOUTH"
	}
	return "interval 1 DAY"
}

func NewRuntimeLog(serviceName string, dateType RuntimeDataType) *RuntimeLog {
	return &RuntimeLog{
		serviceName:     serviceName,
		runtimeDataType: dateType,
	}
}

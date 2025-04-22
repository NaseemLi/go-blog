package common

import (
	"goblog/global"
	"reflect"

	"gorm.io/gorm"
)

type ModelMap interface {
	GetID() uint
}

type ScanOption struct {
	Where *gorm.DB
}

func ScanMap[T ModelMap](model T, option ScanOption) (mp map[uint]T) {
	var list []T
	query := global.DB.Where(model)
	if option.Where != nil {
		query = query.Where(option.Where)
	}
	query.Find(&list)
	mp = make(map[uint]T)
	for _, v := range list {
		mp[v.GetID()] = v
	}
	return
}

func ScanMapV2[T any](model T, option ScanOption) (mp map[uint]T) {
	var list []T
	query := global.DB.Where(model)
	if option.Where != nil {
		query = query.Where(option.Where)
	}
	query.Find(&list)
	mp = make(map[uint]T)
	for _, m := range list {
		v := reflect.ValueOf(m)
		id := v.FieldByName("ID")
		uid, ok := id.Interface().(uint)
		if !ok {
			continue
		}

		mp[uid] = m
	}
	return
}

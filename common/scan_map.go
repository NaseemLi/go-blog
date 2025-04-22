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
	Key   string
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
	key := "ID"
	if option.Key != "" {
		key = option.Key
	}
	for _, m := range list {
		v := reflect.ValueOf(m)
		id := v.FieldByName(key)
		uid, ok := id.Interface().(uint)
		if !ok {
			continue
		}

		mp[uid] = m
	}
	return
}

package redisjwt

import (
	"fmt"
	"goblog/global"
	"goblog/utils/jwts"
	"time"

	"github.com/sirupsen/logrus"
)

type BlackType int8

const (
	UserBlackType   BlackType = 1 //用户注销
	AdminBlackType  BlackType = 2 //管理员手动
	DeviceBlackType BlackType = 3 //其他设备挤下
)

func (b BlackType) String() string {
	return fmt.Sprintf("%d", b)
}

func ParseBlackType(val string) BlackType {
	switch val {
	case "1":
		return UserBlackType
	case "2":
		return AdminBlackType
	case "3":
		return DeviceBlackType
	}
	return UserBlackType
}

func TokenBlack(token string, value BlackType) {
	key := fmt.Sprintf("token_black_%s", token)
	claims, err := jwts.ParseToken(token)
	if err != nil || claims == nil {
		logrus.Errorf("token解析失败 %s", err)
		return
	}

	second := claims.ExpiresAt - time.Now().Unix()

	_, err = global.Redis.Set(key, value.String(), time.Duration(second)*time.Second).Result()
	if err != nil {
		logrus.Errorf("redis 添加黑名单失败 %s", err)
	}
}

func HasTokenBlack(token string) (blackType BlackType, ok bool) {
	key := fmt.Sprintf("token_black_%s", token)
	value, err := global.Redis.Get(key).Result()
	if err != nil {
		return
	}
	blackType = ParseBlackType(value)
	return blackType, true
}

package redisuser

import (
	"goblog/global"
	"strconv"

	"github.com/sirupsen/logrus"
)

type userCacheType string

const (
	userCacheLook userCacheType = "user_look_key"
)

func set(t userCacheType, UserID uint, n int) {
	num, _ := global.Redis.HGet(string(t), strconv.Itoa(int(UserID))).Int()
	num += n
	global.Redis.HSet(string(t), strconv.Itoa(int(UserID)), num)
}

func SetCacheLook(UserID uint, increase bool) {
	var n = 1
	if !increase {
		n = -1
	}
	set(userCacheLook, UserID, n)
}

func get(t userCacheType, UserID uint) int {
	num, _ := global.Redis.HGet(string(t), strconv.Itoa(int(UserID))).Int()
	return num
}

func GetCacheLook(UserID uint) int {
	return get(userCacheLook, UserID)
}

func GetAll(t userCacheType) (mps map[uint]int) {
	result, err := global.Redis.HGetAll(string(t)).Result()
	if err != nil {
		return
	}

	mps = make(map[uint]int)
	for key, numS := range result {
		iK, err := strconv.Atoi(key)
		if err != nil {
			continue
		}
		iN, err := strconv.Atoi(numS)
		if err != nil {
			continue
		}
		mps[uint(iK)] = iN
	}
	return mps
}

func GetAllCacheLook() (mps map[uint]int) {
	return GetAll(userCacheLook)
}

func Clear() {
	err := global.Redis.Del("user_look_key").Err()
	if err != nil {
		logrus.Error(err)
	}
}

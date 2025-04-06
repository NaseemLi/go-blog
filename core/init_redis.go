package core

import (
	"goblog/global"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

func InitRedis() *redis.Client {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr,     // 不写默认就是这个
		Password: global.Config.Redis.Password, // 密码
		DB:       global.Config.Redis.DB,       // 默认是0
	})
	_, err := redisDB.Ping().Result()
	if err != nil {
		logrus.Fatalf("redis 连接失败 %s", err)
	}
	logrus.Info("redis连接成功!")
	return redisDB
}

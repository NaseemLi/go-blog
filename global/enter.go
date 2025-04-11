package global

import (
	"goblog/conf"
	"sync"

	"github.com/go-redis/redis"
	"github.com/mojocn/base64Captcha"
	"gorm.io/gorm"
)

const Version = "0.0.1"

var (
	Config           *conf.Config
	DB               *gorm.DB
	Redis            *redis.Client
	CaptchaStore     = base64Captcha.DefaultMemStore
	EmailVerifyStore = sync.Map{}
)

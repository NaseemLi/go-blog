package global

import (
	"goblog/conf"

	"gorm.io/gorm"
)

var (
	Config *conf.Config
	DB     *gorm.DB
)

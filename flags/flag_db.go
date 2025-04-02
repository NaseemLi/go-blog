package flags

import (
	"goblog/global"
	"goblog/models"

	"github.com/sirupsen/logrus"
)

func FlagDB() {
	err := global.DB.AutoMigrate(
		&models.UserModel{},
		&models.UserConfModel{},
	)
	if err != nil {
		logrus.Errorf("数据库迁移失败 %s", err)
	}

	logrus.Infof("数据库迁移成功!")
}

package core

import (
	"goblog/global"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func InitDB() *gorm.DB {
	if len(global.Config.DB) == 0 {
		logrus.Fatalf("未配置数据库")
	}
	dc := global.Config.DB[0] // 主库（写库）
	db, err := gorm.Open(mysql.Open(dc.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 不生成外键约束
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败: %s", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Fatalf("获取 sql.DB 实例失败: %s", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logrus.Infof("数据库连接成功!")

	if len(global.Config.DB) > 1 {
		// 如果配置了多个数据库，则配置读写分离
		var readList []gorm.Dialector
		for _, d := range global.Config.DB[1:] {
			readList = append(readList, mysql.Open(d.DSN()))
		}

		err := db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dc.DSN())}, // 写库
			Replicas: readList,                               // 读库
			Policy:   dbresolver.RandomPolicy{},              // 随机策略
		}))
		if err != nil {
			logrus.Fatalf("读写配置错误: %s", err)
		}
	}
	return db
}

package core

import (
	"fmt"
	"goblog/global"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func InitDB() *gorm.DB {
	dc := global.Config.DB   //读库
	dc1 := global.Config.DB1 //写库
	db, err := gorm.Open(mysql.Open(dc.DSN()), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //不生成外键约束
	})
	if err != nil {
		logrus.Fatalf("数据库连接失败%s", err)
	}
	fmt.Println(db, err)
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logrus.Infof("数据库连接成功!")

	if !dc1.Empty() {
		//读写库不为空就注册读写分离的配置
		db.Use(dbresolver.Register(dbresolver.Config{
			Sources:  []gorm.Dialector{mysql.Open(dc1.DSN())}, // 写
			Replicas: []gorm.Dialector{mysql.Open(dc.DSN())},  // 读
			Policy:   dbresolver.RandomPolicy{},
		}))

		if err != nil {
			logrus.Fatalf("读写配置错误 %s", err)
		}
	}
	return db
}

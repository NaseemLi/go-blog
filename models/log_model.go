// models/log_model.go
package models

import "goblog/models/enum"

type LogModel struct {
	Model
	LogType     enum.LogType      `json:"logType"` // 日志类型 1 2 3
	Title       string            `gorm:"size:64" json:"title"`
	Content     string            `json:"content"`
	Level       enum.LogLevelType `json:"level"`                      // 日志级别 1 2 3
	UserID      uint              `json:"userID"`                     // 用户id
	UserModel   UserModel         `gorm:"foreignKey:UserID" json:"-"` // 用户信息
	IP          string            `gorm:"size:32" json:"ip"`
	Addr        string            `gorm:"size:64" json:"addr"`
	IsRead      bool              `json:"isRead"` // 是否读取
	LoginStatus bool              `json:"loginStatus"`
	Username    string            `gorm:"size:32" json:"username"`
	Pwd         string            `gorm:"size:32" json:"pwd"`
	LoginType   enum.LoginType    `json:"loginType"`
}

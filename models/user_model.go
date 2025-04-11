package models

import (
	"goblog/models/enum"
	"time"
)

type UserModel struct {
	Model
	Username       string                  `gorm:"size:32" json:"username"`  // 用户名
	Nickname       string                  `gorm:"size:32" json:"nickname"`  // 昵称
	Avatar         string                  `gorm:"size:256" json:"avatar"`   // 头像
	Abstract       string                  `gorm:"size:256" json:"abstract"` // 简介
	RegisterSource enum.RegisterSourceType `json:"registerSource"`           // 注册来源
	CodeAge        int                     `json:"codeAge"`                  // 码龄
	Password       string                  `gorm:"size:64" json:"-"`         // 密码（不序列化为 JSON）
	Email          string                  `gorm:"size:256" json:"email"`    // 邮箱
	OpenID         string                  `gorm:"size:64" json:"openID"`    // 第三方登录的唯一 ID
	Role           enum.RoleType           `json:"role"`                     //1管理员 2用户 3访客
}

type UserConfModel struct {
	UserID             uint       `gorm:"unique" json:"userID"` // 用户ID
	UserModel          UserModel  `gorm:"foreignKey:UserID" json:"-"`
	LikeTags           []string   `gorm:"type:longtext;serializer:json" json:"likeTags"` // 喜欢的标签
	UpdateUsernameDate *time.Time `json:"updateUsernameDate"`                            // 上次修改用户名的时间
	OpenCollect        bool       `json:"openCollect"`                                   // 是否开启收藏
	OpenFollow         bool       `json:"openFollow"`                                    // 是否开启关注
	OpenFans           bool       `json:"openFans"`                                      // 是否开启粉丝
	HomeStyleID        uint       `json:"homeStyleID"`                                   // 首页推荐样式id
}

package models

import (
	"goblog/models/enum"
	"math"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
	UserConfModel  *UserConfModel          `gorm:"foreignKey:UserID" json:"-"`
	Ip             string                  `json:"ip"`
	Addr           string                  `json:"addr"`
	ArticleList    []ArticleModel          `gorm:"foreignKey:UserID" json:"-"` // 文章列表
	LoginList      []UserLoginModel        `gorm:"foreignKey:UserID" json:"-"` // 登录记录
}

func (u UserModel) GetID() uint {
	return u.ID
}

func (u *UserModel) AfterCreate(tx *gorm.DB) error {
	// 创建用户配置记录
	err := tx.Create(&UserConfModel{
		UserID:      u.ID,
		OpenFollow:  true,
		OpenCollect: true,
		OpenFans:    true,
		HomeStyleID: 1,
	}).Error
	if err != nil {
		return err
	}

	// 创建用户消息配置记录
	err = tx.Create(&UserMessageConfModel{
		UserID:             u.ID,
		OpenCommentMessage: true,
		OpenDiggMessage:    true,
		OpenPrivateChat:    true,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserModel) GetCodeAge() int {
	sub := time.Since(u.CreatedAt)
	return int(math.Ceil(sub.Hours() / 24 / 365))
}

func (u *UserModel) BeforeDelete(tx *gorm.DB) (err error) {
	// 定义 model + 表名 的 map，更清晰、更可靠
	models := map[string]interface{}{
		"ArticleDiggModel":            &ArticleDiggModel{},
		"ArticleModel":                &ArticleModel{},
		"CategoryModel":               &CategoryModel{},
		"CollectModel":                &CollectModel{},
		"CommentModel":                &CommentModel{},
		"CommentDiggModel":            &CommentDiggModel{},
		"LogModel":                    &LogModel{},
		"UserArticleCollectModel":     &UserArticleCollectModel{},
		"UserArticleLookHistoryModel": &UserArticleLookHistoryModel{},
		"UserChatActionModel":         &UserChatActionModel{},
		"UserFocusModel":              &UserFocusModel{},
		"UserGlobalNotificationModel": &UserGlobalNotificationModel{},
		"UserLoginModel":              &UserLoginModel{},
		"UserTopArticleModel":         &UserTopArticleModel{},
	}

	for name, model := range models {
		res := tx.Where("user_id = ?", u.ID).Delete(model)
		logrus.Infof("删除关联数据 %s 共 %d 条", name, res.RowsAffected)
	}

	// 删除聊天记录
	var chatList []ChatModel
	tx.Where("send_user_id = ? OR rev_user_id = ?", u.ID, u.ID).Find(&chatList)
	if len(chatList) > 0 {
		tx.Delete(&chatList)
		logrus.Infof("删除关联对话 %d 条", len(chatList))
	}

	// 删除消息（按需）
	var messageList []ChatModel
	tx.Where("rev_user_id = ?", u.ID).Find(&messageList)
	if len(messageList) > 0 {
		tx.Delete(&messageList)
		logrus.Infof("删除关联消息 %d 条", len(messageList))
	}

	return nil
}

type UserConfModel struct {
	UserID             uint       `gorm:"primaryKey;unique" json:"userID"` // 用户ID
	UserModel          UserModel  `gorm:"foreignKey:UserID" json:"-"`
	LikeTags           []string   `gorm:"type:longtext;serializer:json" json:"likeTags"` // 喜欢的标签
	UpdateUsernameDate *time.Time `json:"updateUsernameDate"`                            // 上次修改用户名的时间
	OpenCollect        bool       `json:"openCollect"`                                   // 是否开启收藏
	OpenFollow         bool       `json:"openFollow"`                                    // 是否开启关注
	OpenFans           bool       `json:"openFans"`                                      // 是否开启粉丝
	HomeStyleID        uint       `json:"homeStyleID"`                                   // 首页推荐样式id
	LookCount          int        `json:"lookCount"`                                     // 浏览量
}

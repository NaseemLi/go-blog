package models

import (
	"goblog/models/enum"
	"math"
	"reflect"
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
	var list = []any{
		ArticleDiggModel{},
		ArticleModel{},
		CategoryModel{},
		CollectModel{},
		CommentModel{},
		CommentDiggModel{},
		LogModel{},
		UserArticleCollectModel{},
		UserArticleLookHistoryModel{},
		UserChatActionModel{},
		UserFocusModel{},
		UserGlobalNotificationModel{},
		UserLoginModel{},
		UserTopArticleModel{},
	}
	for _, model := range list {
		count := tx.Delete(model, "user_id = ?", u.ID).RowsAffected
		logrus.Infof("删除关联数据 %T %d条", reflect.TypeOf(model).Name(), count)
	}
	var chatList []ChatModel
	tx.Find("send_user_id = ? OR rev_user_id = ?", u.ID, u.ID).Delete(&chatList)
	logrus.Infof("删除关联对话%d条", len(chatList))

	var messageList []ChatModel
	tx.Find("rev_user_id = ?", u.ID).Delete(&messageList)
	logrus.Infof("删除关联消息%d条", len(messageList))

	return
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

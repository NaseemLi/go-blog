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
		&models.ArticleModel{},
		&models.CategoryModel{},
		&models.ArticleDiggModel{},
		&models.BannerModel{},
		&models.CollectModel{},
		&models.CommentModel{},
		&models.GlobalNotificationModel{},
		&models.ImageModel{},
		&models.Model{},
		&models.UserLoginModel{},
		&models.UserTopArticleModel{},
		&models.UserArticleLookHistoryModel{},
		&models.LogModel{},
		&models.UserArticleCollectModel{},
		&models.CommentDiggModel{},
		&models.MessageModel{},
		&models.UserMessageConfModel{},
		&models.UserGlobalNotificationModel{},
		&models.UserFocusModel{},
		&models.ChatModel{},
		&models.UserChatActionModel{},
	)
	if err != nil {
		logrus.Errorf("数据库迁移失败 %s", err)
	}

	logrus.Infof("数据库迁移成功!")
}

package cronservice

import (
	"goblog/global"
	"goblog/models"
	redisarticle "goblog/service/redis_service/redis_article"
	redisuser "goblog/service/redis_service/redis_user"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SyncUser() {
	lookMap := redisarticle.GetAllCacheLook()

	var list []models.ArticleModel
	global.DB.Find(&list)

	for _, model := range list {
		look := lookMap[model.ID]
		if look == 0 {
			continue
		}

		err := global.DB.Model(&model).Updates(map[string]any{
			"look_count": gorm.Expr("look_count + ?", look),
		}).Error

		if err != nil {
			logrus.Errorf("更新失败: %s", err)
			return
		}
		logrus.Infof("%s 更新成功", model.UserID)
	}
	//清除缓存
	redisuser.Clear()
}

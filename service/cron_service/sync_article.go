package cronservice

import (
	"goblog/global"
	"goblog/models"
	redisarticle "goblog/service/redis_service/redis_article"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SyncArticle() {
	//todo:大数据量加锁
	collectMap := redisarticle.GetAllCacheCollect()
	diggMap := redisarticle.GetAllCacheDigg()
	lookMap := redisarticle.GetAllCacheLook()

	var list []models.ArticleModel
	global.DB.Find(&list)

	for _, model := range list {
		collect := collectMap[model.ID]
		digg := diggMap[model.ID]
		look := lookMap[model.ID]
		if collect == 0 || digg == 0 || look == 0 {
			continue
		}

		err := global.DB.Model(&model).Updates(map[string]any{
			"look_count":    gorm.Expr("look_count + ?", look),
			"digg_count":    gorm.Expr("digg_count + ?", digg),
			"collect_count": gorm.Expr("collect_count + ?", collect),
		}).Error

		if err != nil {
			logrus.Errorf("更新失败: %s", err)
			return
		}
		logrus.Infof("%s更新成功", model.Title)
	}
	//todo:这里可能有增量数据
	//todo:可以再获取一次
	// collectMap := redisarticle.GetAllCacheCollect()
	// diggMap := redisarticle.GetAllCacheDigg()
	// lookMap := redisarticle.GetAllCacheLook()

	//清除缓存
	redisarticle.Clear()
}

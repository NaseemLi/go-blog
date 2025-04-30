package cronservice

import (
	"goblog/global"
	"goblog/models"
	redissite "goblog/service/redis_service/redis_site"
)

func SyncSiteFlow() {
	flow := redissite.GetFlow()
	global.DB.Create(&models.SiteFlowModel{Count: flow})

	redissite.ClearFlow()
}

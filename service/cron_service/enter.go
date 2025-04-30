package cronservice

import (
	"time"

	"github.com/robfig/cron/v3"
)

func Cron() {
	//crontab := cron.New()  默认从分开始进行时间调度
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	crontab := cron.New(cron.WithSeconds(), cron.WithLocation(timezone))
	//每天两点同步文章数据
	crontab.AddFunc("0 0 2 * * *", SyncArticle)
	crontab.AddFunc("0 30 2 * * *", SyncUser)
	crontab.AddFunc("0 0 2 * * *", SyncComment)
	crontab.AddFunc("0 59 23 * * *", SyncSiteFlow)

	crontab.Start()
}

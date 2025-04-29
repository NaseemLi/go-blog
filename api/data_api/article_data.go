package dataapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	"goblog/models/enum"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleDataResponse struct {
	GrowthRate int      `json:"growthRate"`
	GrowthNum  int      `json:"growthNum"`
	DateList   []string `json:"dateList"`
	CountList  []int    `json:"countList"`
}

func (DataApi) ArticleDataView(c *gin.Context) {
	now := time.Now()
	startDate := now.AddDate(0, 0, -6) // 最近7天，包括今天

	var articleList []models.ArticleModel
	err := global.DB.
		Where("created_at BETWEEN ? AND ? AND status = ?", startDate, now, enum.ArticleStatusPublished).
		Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		res.FailWithError(err, c)
		return
	}

	// 按天统计数量
	dateMap := make(map[string]int, 7)
	for _, article := range articleList {
		dateStr := article.CreatedAt.Format("2006-01-02")
		dateMap[dateStr]++
	}

	// 构造返回数据
	response := ArticleDataResponse{}
	for i := 0; i < 7; i++ {
		date := startDate.AddDate(0, 0, i)
		dateStr := date.Format("2006-01-02")
		count := dateMap[dateStr]
		response.DateList = append(response.DateList, dateStr)
		response.CountList = append(response.CountList, count)
	}

	// 计算增长量
	if len(response.CountList) >= 7 {
		response.GrowthNum = response.CountList[6] - response.CountList[5]

		if response.CountList[5] == 0 {
			if response.CountList[6] > 0 {
				response.GrowthRate = 100 // 昨天 0 篇，今天有增长，设为100%
			} else {
				response.GrowthRate = 0 // 都没有增长
			}
		} else {
			response.GrowthRate = int(float64(response.GrowthNum) / float64(response.CountList[5]) * 100)
		}
	}

	res.OkWithData(response, c)
}

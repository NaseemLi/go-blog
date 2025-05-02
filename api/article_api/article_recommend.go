package articleapi

import (
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"

	"github.com/gin-gonic/gin"
)

type ArticleRecommentResponse struct {
	ID        uint   `json:"id" gorm:"column:id"`
	Title     string `json:"title" gorm:"column:title"`
	LookCount int    `json:"lookCount" gorm:"column:lookCount" `
}

func (ArticleApi) ArticleRecommentView(c *gin.Context) {
	cr := middleware.GetBind[common.PageInfo](c)

	var list = make([]ArticleRecommentResponse, 0)
	global.DB.Model(models.ArticleModel{}).
		Order("look_count desc").
		Where("date(created_at) = date(now())").
		Limit(cr.Limit).
		Select("id, title, look_count").
		Scan(&list)

	res.OkWithList(list, len(list), c)
}

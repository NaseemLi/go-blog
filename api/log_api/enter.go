package logapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/models"

	"github.com/gin-gonic/gin"
)

type LogApi struct {
}

type LogListRequest struct {
	Limit int    `form:"limit"`
	Page  int    `form:"page"`
	Key   string `form:"key"`
}

func (LogApi) LogListView(c *gin.Context) {
	//分页查询 精确查询,模糊匹配
	var cr LogListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var list []models.LogModel
	if cr.Page > 20 {
		cr.Page = 1
	}
	if cr.Page <= 0 {
		cr.Page = 1
	}
	if cr.Limit == 0 || cr.Limit > 100 {
		cr.Limit = 10
	}

	offset := (cr.Page - 1) * cr.Limit
	global.DB.Debug().Offset(offset).Limit(cr.Limit).Find(&list)

	var count int64
	global.DB.Debug().Model(models.LogModel{}).Count(&count)

	res.OkWithList(list, int(count), c)
}

package articleapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"

	"github.com/gin-gonic/gin"
)

// ArticleStatusExamineRequest 审核文章请求
type ArticleExamineRequest struct {
	ArticleID uint   `json:"articleID" binding:"required"`
	Status    int    `json:"status" binding:"required,oneof=3 4"`
	Msg       string `json:"msg"` //为4的时候,传递进入
}

func (ArticleApi) ArticleExamineView(c *gin.Context) {
	cr := middleware.GetBind[ArticleExamineRequest](c)

	var article models.ArticleModel
	err := global.DB.Take(&article, cr.ArticleID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}
	global.DB.Model(&article).Update("status", cr.Status)

	// todo发送系统通知

	res.OkWithMsg("文章审核成功", c)
}

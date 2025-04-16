package articleapi

import (
	"fmt"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"

	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleRemoveAdminView(c *gin.Context) {
	cr := middleware.GetBind[models.RemoveRequest](c)

	var list []models.ArticleModel
	global.DB.Where("id IN ?", cr.IDList).Find(&list)

	if len(list) > 0 {
		err := global.DB.Delete(&list).Error
		if err != nil {
			res.FailWithMsg("删除文章失败", c)
			return
		}
	}

	res.OkWithMsg(fmt.Sprintf("删除成功,删除成功%d条", len(list)), c)
}

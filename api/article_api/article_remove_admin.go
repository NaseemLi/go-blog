package articleapi

import (
	"fmt"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	messageservice "goblog/service/message_service"

	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleRemoveAdminView(c *gin.Context) {
	cr := middleware.GetBind[models.RemoveRequest](c)

	var list []models.ArticleModel
	global.DB.Where("id IN ?", cr.IDList).Find(&list)

	if len(list) > 0 {
		for _, model := range list {
			messageservice.InsertSystemMessage(model.UserID, "管理员删除了你的文章", fmt.Sprintf("%s内容不符合社区规范", model.Title), "", "")
		}
		err := global.DB.Delete(&list).Error
		if err != nil {
			res.FailWithMsg("删除文章失败", c)
			return
		}
	}

	res.OkWithMsg(fmt.Sprintf("删除成功,删除成功%d条", len(list)), c)
}

package articleapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleRemoveUserView(c *gin.Context) {
	cr := middleware.GetBind[models.IDRequest](c)

	claims := jwts.GetClaims(c)
	var model models.ArticleModel
	err := global.DB.Take(&model, "user_id = ? AND id = ?", claims.UserID, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	err = global.DB.Delete(&model).Error
	if err != nil {
		res.FailWithMsg("删除文章失败", c)
		return
	}
	res.OkWithMsg("删除成功", c)
}

package articleapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	"goblog/models/ctype"
	"goblog/models/enum"
	"goblog/utils"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleTagOptions(c *gin.Context) {
	claims := jwts.GetClaims(c)

	var articleList []models.ArticleModel
	global.DB.Find(&articleList, "user_id = ? and status = ?", claims.UserID, enum.ArticleStatusPublished)

	var tagList ctype.List
	for _, v := range articleList {
		tagList = append(tagList, v.TagList...)
	}

	tagList = utils.Unique(tagList)
	var list = make([]models.OptionsResponse[string], 0)

	for _, v := range tagList {
		list = append(list, models.OptionsResponse[string]{
			Value: v,
			Label: v,
		})
	}
	res.OkWithData(list, c)
}

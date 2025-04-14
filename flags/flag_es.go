package flags

import (
	"goblog/models"
	esservice "goblog/service/es_service"
)

func Esindex() {
	article := models.ArticleModel{}
	esservice.CreateIndexV2(article.Index(), article.Mapping())
}

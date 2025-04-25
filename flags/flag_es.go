package flags

import (
	"goblog/global"
	"goblog/models"
	esservice "goblog/service/es_service"

	"github.com/sirupsen/logrus"
)

func Esindex() {
	if global.ESClient == nil {
		logrus.Warnf("es client 为空")
		return
	}
	article := models.ArticleModel{}
	esservice.CreateIndexV2(article.Index(), article.Mapping())
	text := models.TextModel{}
	esservice.CreateIndexV2(text.Index(), text.Mapping())
}

package commentapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	commentservice "goblog/service/comment_service"

	"github.com/gin-gonic/gin"
)

func (CommentApi) CommentTreeView(c *gin.Context) {
	var cr = middleware.GetBind[models.IDRequest](c)

	var article models.ArticleModel
	err := global.DB.Take(&article, "id = ? and status = ?", cr.ID, enum.ArticleStatusPublished).Error
	if err != nil {
		res.FailWithMsg("文章不存在或已被删除", c)
		return
	}

	//查出根评论
	var commentList []models.CommentModel
	global.DB.Find(&commentList, "article_id = ? and parent_id is null", cr.ID)
	var list = make([]commentservice.CommentResponse, 0)
	for _, model := range commentList {
		response := commentservice.GetCommentTreeV3(model.ID)
		list = append(list, *response)
	}

	res.OkWithList(list, len(list), c)
}

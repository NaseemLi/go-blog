package commentapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	commentservice "goblog/service/comment_service"
	messageservice "goblog/service/message_service"
	redisarticle "goblog/service/redis_service/redis_article"
	rediscomment "goblog/service/redis_service/redis_comment"

	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type CommentCreateRequest struct {
	Content   string `json:"content" binding:"required"`
	ArticleID uint   `json:"articleID" binding:"required"`
	ParentID  *uint  `json:"parentID"` // 父评论
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	cr := middleware.GetBind[CommentCreateRequest](c)

	var article models.ArticleModel
	err := global.DB.Take(&article, "id = ? and status = ?", cr.ArticleID, enum.ArticleStatusPublished).Error
	if err != nil {
		res.FailWithMsg("文章不存在或已被删除", c)
		return
	}

	claims := jwts.GetClaims(c)

	model := models.CommentModel{
		Content:   cr.Content,
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
		ParentID:  cr.ParentID,
	}

	//去找这个评论的根评论
	if cr.ParentID != nil {
		//查找父评论
		parentList := commentservice.GetParents(*cr.ParentID)
		//判断层级是否满足
		if len(parentList)+1 > global.Config.Site.Article.CommentLine {
			res.FailWithMsg("评论层级超过限制", c)
			return
		}
		if len(parentList) > 0 {
			model.RootParentID = &parentList[len(parentList)-1].ID
			for _, v := range parentList {
				rediscomment.SetCacheApply(v.ID, 1)
			}
			//给父评论的拥有人发送消息
			defer func() {
				go messageservice.InsertApplyMessage(model)
			}()
		}
	}

	err = global.DB.Create(&model).Error
	if err != nil {
		res.FailWithMsg("评论失败", c)
		return
	}

	go messageservice.InsertCommentMessage(model)
	redisarticle.SetCacheComment(article.ID, 1)

	res.OkWithMsg("评论成功", c)
}

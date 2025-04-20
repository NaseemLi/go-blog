package articleapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	messageservice "goblog/service/message_service"
	redisarticle "goblog/service/redis_service/redis_article"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleDiggView(c *gin.Context) {
	cr := middleware.GetBind[models.IDRequest](c)

	var article models.ArticleModel
	err := global.DB.Take(&article, "status = ? and id = ?", enum.ArticleStatusPublished, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	claims := jwts.GetClaims(c)
	//查看之前有没有点过赞
	var userDiggArticle models.ArticleDiggModel
	err = global.DB.Take(&userDiggArticle, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	if err != nil {
		//点赞
		model := models.ArticleDiggModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		}
		err = global.DB.Create(&model).Error
		if err != nil {
			res.FailWithMsg("点赞失败", c)
			return
		}
		redisarticle.SetCacheDigg(cr.ID, true)
		//给文章拥有者发消息
		go messageservice.InsertDiggArticleMessage(model)
		res.OkWithMsg("点赞成功", c)
		return
	}
	//取消点赞
	global.DB.Delete(&userDiggArticle, "user_id = ? and article_id = ?", claims.UserID, cr.ID)
	res.OkWithMsg("取消点赞成功", c)
	redisarticle.SetCacheDigg(cr.ID, false)
}

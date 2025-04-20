package articleapi

import (
	"fmt"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	messageservice "goblog/service/message_service"
	redisarticle "goblog/service/redis_service/redis_article"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ArticleCollectRequest struct {
	ArticleID uint `json:"articleID" binding:"required"`
	CollectID uint `json:"collectID"`
}

func (ArticleApi) ArticleCollectView(c *gin.Context) {
	cr := middleware.GetBind[ArticleCollectRequest](c)

	var article models.ArticleModel
	err := global.DB.Take(&article, "status = ? and id = ?", enum.ArticleStatusPublished, cr.ArticleID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}
	var collectModel models.CollectModel
	claims := jwts.GetClaims(c)
	if cr.CollectID == 0 {
		// 是默认收藏夹
		err = global.DB.Take(&collectModel, "user_id = ? and is_default = ?", claims.UserID, 1).Error
		if err != nil {
			// 创建一个默认收藏夹
			collectModel.Title = "默认收藏夹"
			collectModel.UserID = claims.UserID
			collectModel.IsDefault = true
			global.DB.Create(&collectModel)
		}
		cr.CollectID = collectModel.ID
	} else {
		// 判断收藏夹是否存在，并且是否是自己创建的
		err = global.DB.Take(&collectModel, "user_id = ? ", claims.UserID).Error
		if err != nil {
			res.FailWithMsg("收藏夹不存在", c)
			return
		}
	}

	// 判断是否收藏
	var articleCollect models.UserArticleCollectModel
	err = global.DB.Where(models.UserArticleCollectModel{
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
		CollectID: cr.CollectID,
	}).Take(&articleCollect).Error

	if err != nil {
		// 收藏
		model := models.UserArticleCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ArticleID,
			CollectID: cr.CollectID,
		}
		err = global.DB.Create(&model).Error
		if err != nil {
			res.FailWithMsg("收藏失败", c)
			return
		}

		res.OkWithMsg("收藏成功", c)

		// 对收藏夹进行加1
		redisarticle.SetCacheCollect(cr.CollectID, true)
		go messageservice.InsertCollectArticleMessage(model)
		global.DB.Model(&collectModel).Update("article_count", gorm.Expr("article_count + 1"))
		return
	}
	// 取消收藏
	err = global.DB.Where(models.UserArticleCollectModel{
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
		CollectID: cr.CollectID,
	}).Delete(&models.UserArticleCollectModel{}).Error

	if err != nil {
		res.FailWithMsg("取消收藏失败", c)
		return
	}
	res.OkWithMsg("取消收藏成功", c)
	//TODO:收藏数同步缓存
	redisarticle.SetCacheCollect(cr.CollectID, false)
	global.DB.Model(&collectModel).Update("article_count", gorm.Expr("article_count - 1"))
}

func (ArticleApi) ArticleCollectPatchRemoveView(c *gin.Context) {
	var cr = middleware.GetBind[models.RemoveRequest](c)

	claims := jwts.GetClaims(c)
	var userCollectList []models.UserArticleCollectModel
	global.DB.Find(&userCollectList, "id in ? and user_id = ?", cr.IDList, claims.UserID)

	if len(userCollectList) > 0 {
		global.DB.Delete(&userCollectList)
	}
	res.OkWithMsg(fmt.Sprintf("成功删除 %d 篇", len(userCollectList)), c)
}

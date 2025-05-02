package articleapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	redisarticle "goblog/service/redis_service/redis_article"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type ArticleDetailResponse struct {
	models.ArticleModel
	Username      string  `json:"username"`
	Nickname      string  `json:"nickname"`
	UserAvatar    string  `json:"userAvatar"`
	CategoryTitle *string `json:"categoryTitle"`
	IsDigg        bool    `json:"isDigg"`    //是否点赞
	IsCollect     bool    `json:"isCollect"` //是否收藏
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	cr := middleware.GetBind[models.IDRequest](c)

	//未登录的用户,只能看到发布成功的文章

	//登录用户,能看到自己的所有文章

	//管理员,能看到全部文章

	var article models.ArticleModel
	err := global.DB.Preload("UserModel").Preload("CategoryModel").Take(&article, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		//未登录的
		if article.Status != enum.ArticleStatusPublished {
			res.FailWithMsg("文章不存在", c)
			return
		}
	}

	data := ArticleDetailResponse{
		ArticleModel: article,
		Username:     article.UserModel.Username,
		Nickname:     article.UserModel.Nickname,
		UserAvatar:   article.UserModel.Avatar,
	}

	if err == nil && claims != nil {
		switch claims.Role {
		case enum.UserRole:
			if claims.UserID == article.UserID {
				//登录的人看的不是自己的
				if article.Status != enum.ArticleStatusPublished {
					res.FailWithMsg("文章不存在", c)
					return
				}
			}
		}
		//查用户是否收藏了文章,点赞了文章
		var userDiggModel models.ArticleDiggModel
		err = global.DB.Take(&userDiggModel, "user_id = ? and article_id = ?", claims.UserID, article.ID).Error
		if err == nil {
			data.IsDigg = true
		}

		//用户是否收藏了文章 只要在其中任何一个收藏夹内 都判断用户收藏了文章
		var userCollectModel models.UserArticleCollectModel
		err = global.DB.Take(&userCollectModel, "user_id = ? and article_id = ?", claims.UserID, article.ID).Error
		if err == nil {
			data.IsCollect = true
		}
	}

	lookCount := redisarticle.GetCacheLook(article.ID)
	diggCount := redisarticle.GetCacheDigg(article.ID)
	collectCount := redisarticle.GetCacheCollect(article.ID)
	commentCount := redisarticle.GetCacheComment(article.ID)

	data.DiggCount = article.DiggCount + diggCount
	data.CollectCount = article.CollectCount + collectCount
	data.LookCount = article.LookCount + lookCount
	data.CommentCount = article.CommentCount + commentCount

	if article.CategoryModel != nil {
		data.CategoryTitle = &article.CategoryModel.Title
	}

	res.OkWithData(data, c)
}

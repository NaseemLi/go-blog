package userapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type UserArticleTopRequest struct {
	ArticleID uint `json:"articleID" binding:"required"`
	Type      int8 `json:"type" binding:"required,oneof=1 2"`
}

func (UserApi) UserArticleTopView(c *gin.Context) {
	cr := middleware.GetBind[UserArticleTopRequest](c)

	var model models.ArticleModel
	err := global.DB.Take(&model, cr.ArticleID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	claims := jwts.GetClaims(c)

	switch cr.Type {
	case 1:
		//用户置顶文章
		//验证文章是不是自己的,并且是已发布的
		if model.UserID != claims.UserID {
			res.FailWithMsg("只能置顶自己的文章", c)
			return
		}

		if model.Status != enum.ArticleStatusPublished {
			res.FailWithMsg("只能置顶已发布的文章", c)
			return
		}
		//判断之前自己是否置顶过
		var userTopArticleList []models.UserTopArticleModel
		global.DB.Find(&userTopArticleList, "user_id = ? and article_id = ?", claims.UserID, cr.ArticleID)

		if len(userTopArticleList) == 0 {
			//置顶
			global.DB.Create(&models.UserTopArticleModel{
				UserID:    claims.UserID,
				ArticleID: cr.ArticleID,
			})
			res.OkWithMsg("置顶文章成功", c)
			return
		}

		if len(userTopArticleList) == 1 {
			uta := userTopArticleList[0]
			if uta.ArticleID != cr.ArticleID {
				res.FailWithMsg("普通用户只能置顶一片文章", c)
				return
			}
		}

		uta := userTopArticleList[0]
		global.DB.Delete(&uta)
		res.OkWithMsg("取消置顶成功", c)
		return

		//查不到,自己从来没有置顶过
		//查到=1 判断是不是这篇 是的话取消置顶 不是的话 超过了最大置顶限度

	case 2:
		//管理员置顶文章
		if claims.Role != enum.AdminRole {
			res.FailWithMsg("没有权限置顶文章", c)
			return
		}

		if model.Status != enum.ArticleStatusPublished {
			res.FailWithMsg("只能置顶已发布的文章", c)
			return
		}

		var userTopArticle []models.UserTopArticleModel
		err := global.DB.Take(&userTopArticle, "article_id = ? and user_id = ?", cr.ArticleID, claims.UserID).Error
		if err != nil {
			global.DB.Create(&models.UserTopArticleModel{
				UserID:    claims.UserID,
				ArticleID: cr.ArticleID,
			})
			res.OkWithMsg("置顶文章成功", c)
			return
		}
		global.DB.Delete(&userTopArticle)
		res.OkWithMsg("取消置顶成功", c)
		return
	}
}

package articleapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/ctype"
	"goblog/models/enum"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type ArticleCreateRequest struct {
	Title       string                 `json:"title" binding:"required"`
	Abstract    string                 `json:"abstract"`
	Content     string                 `json:"content" binding:"required"`
	CategoryID  *uint                  `json:"categoryID"`
	TagList     ctype.List             `json:"tagList"`
	Cover       string                 `json:"cover"`
	OpenComment bool                   `json:"openComment"`
	Status      enum.ArticleStatusType `json:"status" binding:"required,oneof = 1 2"`
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	cr := middleware.GetBind[ArticleCreateRequest](c)

	user, err := jwts.GetClaims(c).GetUser()
	if err != nil {
		res.FailWithMsg("用户不存在", c)
	}

	//判断分类 ID 是不是自己创建的

	//文章正文防 xss 注入

	//正文内容图片转存

	//AI 分析
	var article = models.ArticleModel{
		Title:       cr.Title,
		Abstract:    cr.Abstract,
		Content:     cr.Content,
		UserID:      user.ID,
		TagList:     cr.TagList,
		Cover:       cr.Cover,
		OpenComment: cr.OpenComment,
		CategoryID:  cr.CategoryID,
		Status:      cr.Status,
	}
	if global.Config.Site.Article.NoExamine {
		article.Status = enum.ArticleStatusPublished
	}
	err = global.DB.Create(&article).Error
	if err != nil {
		res.FailWithMsg("文章创建失败", c)
		return
	}
	res.OkWithMsg("文章创建成功", c)
}

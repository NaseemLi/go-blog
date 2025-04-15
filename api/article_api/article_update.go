package articleapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/ctype"
	"goblog/models/enum"
	"goblog/utils/jwts"
	"goblog/utils/markdown"
	"goblog/utils/xss"

	"github.com/gin-gonic/gin"
)

type ArticleUpdateRequest struct {
	ID          uint       `json:"id" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Abstract    string     `json:"abstract"`
	Content     string     `json:"content" binding:"required"`
	CategoryID  *uint      `json:"categoryID"`
	TagList     ctype.List `json:"tagList"`
	Cover       string     `json:"cover"`
	OpenComment bool       `json:"openComment"`
}

func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	cr := middleware.GetBind[ArticleUpdateRequest](c)
	claims := jwts.GetClaims(c)
	if claims == nil {
		res.FailWithMsg("未登录或 token 错误", c)
		return
	}

	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	var article models.ArticleModel
	err = global.DB.Take(&article, cr.ID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	// 更新的文章必须是自己的
	if article.UserID != user.ID {
		res.FailWithMsg("只能更新自己的文章", c)
		return
	}

	// 判断分类id是不是自己创建的
	var category models.CategoryModel
	if cr.CategoryID != nil {
		err = global.DB.Take(&category, "id = ? and user_id = ?", *cr.CategoryID, user.ID).Error
		if err != nil {
			res.FailWithMsg("文章分类不存在", c)
			return
		}
	}

	// 文章正文防xss注入
	cr.Content = xss.XSSFilter(cr.Content)

	// 如果不传简介，那么从正文中取前30个字符
	if cr.Abstract == "" {
		abs, err1 := markdown.ExtractContent(cr.Content, 100)
		if err1 != nil {
			res.FailWithMsg("正文解析错误", c)
			return
		}
		cr.Abstract = abs
	}
	mps := map[string]any{
		"title":        cr.Title,
		"abstract":     cr.Abstract,
		"content":      cr.Content,
		"category_id":  cr.CategoryID,
		"tag_list":     cr.TagList,
		"cover":        cr.Cover,
		"open_comment": cr.OpenComment,
	}
	if article.Status == enum.ArticleStatusPublished && !global.Config.Site.Article.NoExamine {
		// 如果是已发布的文章，进行编辑，那么就要改成待审核
		mps["status"] = enum.ArticleStatusExamine
	}

	err = global.DB.Model(&article).Updates(mps).Error
	if err != nil {
		res.FailWithMsg("更新失败", c)
		return
	}

	res.OkWithMsg("文章更新成功", c)
}

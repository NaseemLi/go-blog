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

type ArticleCreateRequest struct {
	Title       string     `json:"title" binding:"required"`
	Abstract    string     `json:"abstract"`
	Content     string     `json:"content" binding:"required"`
	CategoryID  *uint      `json:"categoryID"`
	TagList     ctype.List `json:"tagList"`
	Cover       string     `json:"cover"`
	OpenComment bool       `json:"openComment"`
	Status      int        `json:"status" binding:"required,oneof=1 2"`
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	cr := middleware.GetBind[ArticleCreateRequest](c)

	user, err := jwts.GetClaims(c).GetUser()
	if err != nil {
		res.FailWithMsg("用户不存在", c)
	}

	//判断分类 ID 是不是自己创建的
	var category models.CategoryModel
	if cr.CategoryID != nil {
		err = global.DB.Take(&category, "id = ? and user_id = ?", &cr.CategoryID, user.ID).Error
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
		Status:      enum.ArticleStatusType(cr.Status), // 显式转换成你定义的枚举类型
	}
	//免审核
	if cr.Status == int(enum.ArticleStatusExamine) && global.Config.Site.Article.NoExamine {
		article.Status = enum.ArticleStatusPublished
	}
	err = global.DB.Create(&article).Error
	if err != nil {
		res.FailWithMsg("文章创建失败", c)
		return
	}
	res.OkWithMsg("文章创建成功", c)
}

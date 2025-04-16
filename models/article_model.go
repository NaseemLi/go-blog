package models

import (
	_ "embed"
	"goblog/global"
	"goblog/models/ctype"
	"goblog/models/enum"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleModel struct {
	Model
	Title        string                 `gorm:"size:32" json:"title"`
	Abstract     string                 `gorm:"size:256" json:"abstract"`
	Content      string                 `json:"content,omitempty"`
	CategoryID   *uint                  `json:"categoryID"`                   // 分类的id
	TagList      ctype.List             `gorm:"type:longtext" json:"tagList"` // 标签列表
	Cover        string                 `gorm:"size:256" json:"cover"`
	UserID       uint                   `json:"userID"`
	UserModel    UserModel              `gorm:"foreignKey:UserID" json:"-"`
	LookCount    int                    `json:"lookCount"`
	DiggCount    int                    `json:"diggCount"`
	CommentCount int                    `json:"commentCount"`
	CollectCount int                    `json:"collectCount"`
	OpenComment  bool                   `json:"openComment"` // 开启评论
	Status       enum.ArticleStatusType `json:"status"`      // 状态 草稿 审核中  已发布
}

//go:embed mappings/article_mapping.json
var articleMapping string

func (ArticleModel) Mapping() string {
	return articleMapping
}

func (ArticleModel) Index() string {
	return "article_index"
}

func (a *ArticleModel) BeforeDelete(tx *gorm.DB) (err error) {
	//评论
	var commentList []CommentModel
	global.DB.Find(&commentList, "article_id = ?", a.ID).Delete(&commentList)
	//点赞
	var diggList []ArticleDiggModel
	global.DB.Where("article_id = ?", a.ID).Delete(&diggList)
	//收藏
	var collectList []CollectModel
	global.DB.Where("article_id = ?", a.ID).Delete(&collectList)
	//置顶
	var topList []UserTopArticleModel
	global.DB.Where("article_id = ?", a.ID).Delete(&topList)
	//浏览
	var lookList []UserArticleLookHistoryModel
	global.DB.Where("article_id = ?", a.ID).Delete(&lookList)

	logrus.Infof("删除关联评论%d条", len(commentList))
	logrus.Infof("删除关联点赞%d条", len(diggList))
	logrus.Infof("删除关联收藏%d条", len(collectList))
	logrus.Infof("删除关联置顶%d条", len(topList))
	logrus.Infof("删除关联浏览记录%d条", len(lookList))

	return
}

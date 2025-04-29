// models/user_top_article_model.go
package models

import "time"

type UserTopArticleModel struct {
	Model
	UserID       uint         `gorm:"uniqueIndex:idx_name" json:"userID"`
	ArticleID    uint         `gorm:"uniqueIndex:idx_name" json:"articleID"`
	UserModel    UserModel    `gorm:"foreignKey:UserID" json:"-"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID" json:"-"`
	CreatedAt    time.Time    `json:"createdAt"` // 收藏的时间
}

// models/collect_model.go
package models

type CollectModel struct {
	Model
	Title       string                    `gorm:"size:32" json:"title"`
	Abstract    string                    `gorm:"size:256" json:"abstract"`
	Cover       string                    `gorm:"size:256" json:"cover"`
	ArticleList []UserArticleCollectModel `gorm:"foreignKey:CollectID" json:"-"`
	UserID      uint                      `json:"userID"`
	IsDefault   bool                      `json:"isDefault"` //是否默认收藏夹
	UserModel   UserModel                 `gorm:"foreignKey:UserID" json:"-"`
}

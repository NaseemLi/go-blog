package commentapi

import (
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	"goblog/utils/jwts"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentListRequest struct {
	common.PageInfo
	ArticleID uint `form:"articleID"`
	UserID    uint `form:"userID"`
	Type      int  `form:"type" binding:"required"` // 1.查我发文章的评论 2.查我发布的评论 3.管理员所有的评论
}

type CommentListResponse struct {
	ID           uint      `json:"id"`
	CreateAt     time.Time `json:"createAt"`
	Content      string    `json:"content"`
	UserID       uint      `json:"userID"`
	UserNickname string    `json:"userNickname"`
	UserAvatar   string    `json:"userAvatar"`
	ArticleID    uint      `json:"articleID"`
	ArticleTitle string    `json:"articleTitle"`
	ArticleCover string    `json:"articleCover"`
	DiggCount    int       `json:"diggCount"`
}

func (CommentApi) CommentListView(c *gin.Context) {
	cr := middleware.GetBind[CommentListRequest](c)
	query := global.DB.Where("")
	claims := jwts.GetClaims(c)

	switch cr.Type {
	case 1: //查我发了哪些文章
		var articleIDList []uint
		global.DB.Model(&models.ArticleModel{}).Where("user_id = ? and status = ?", claims.UserID, enum.ArticleStatusPublished).Select("id").Scan(&articleIDList)
		query.Where("article_id in ?", articleIDList)
		cr.UserID = 0
	case 2: //查我发布的评论
		cr.UserID = claims.UserID
	case 3:
	}

	_list, _, _ := common.ListQuery(models.CommentModel{
		ArticleID: cr.ArticleID,
		UserID:    cr.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Likes:    []string{"content"},
		Preloads: []string{"UserModel", "ArticleModel"},
		Where:    query,
	})

	list := make([]CommentListResponse, 0)
	for _, model := range _list {
		list = append(list, CommentListResponse{
			ID:           model.ID,
			CreateAt:     model.CreatedAt,
			Content:      model.Content,
			UserID:       model.UserID,
			UserNickname: model.UserModel.Nickname,
			UserAvatar:   model.UserModel.Avatar,
			ArticleID:    model.ArticleID,
			ArticleTitle: model.ArticleModel.Title,
			ArticleCover: model.ArticleModel.Cover,
			DiggCount:    model.DiggCount,
		})
	}
	res.OkWithList(list, len(list), c)
}

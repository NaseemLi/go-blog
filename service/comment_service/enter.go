package commentservice

import (
	"goblog/global"
	"goblog/models"
	"time"
)

func GetRootComment(commentID uint) (model *models.CommentModel) {
	var comment models.CommentModel
	err := global.DB.Take(&comment, commentID).Error
	if err != nil {
		return nil
	}
	if comment.ParentID == nil {
		//没有父评论了,那么他就是根评论
		return &comment
	}
	return GetRootComment(*comment.ParentID)
}

// GetCommentTree 获取评论树
func GetCommentTree(model *models.CommentModel) {
	global.DB.Preload("SubCommentList").Take(model)
	for _, commentModel := range model.SubCommentList {
		GetCommentTree(commentModel)
	}
}

// GetCommentTreeV2 获取评论树
func GetCommentTreeV2(id uint) (model *models.CommentModel) {
	model = &models.CommentModel{
		Model: models.Model{ID: id},
	}

	global.DB.Preload("SubCommentList").Take(model)
	for i := 0; i < len(model.SubCommentList); i++ {
		commentModel := model.SubCommentList[i]
		item := GetCommentTreeV2(commentModel.ID)
		model.SubCommentList[i] = item
	}
	return
}

type CommentResponse struct {
	ID           uint               `json:"id"`
	CreatedAt    time.Time          `json:"createdAt"`
	Content      string             `json:"content"`
	UserID       uint               `json:"userID"`
	UserNickname string             `json:"userNickname"`
	UserAvatar   string             `json:"userAvatar"`
	ArticleID    uint               `json:"articleID"`
	ParentID     *uint              `json:"parentID"`
	DiggCount    int                `json:"diggCount"`
	ApplyCount   int                `json:"applyCount"`
	SubComments  []*CommentResponse `json:"subComments"`
}

func GetCommentTreeV3(id uint) *CommentResponse {
	model := &models.CommentModel{}
	err := global.DB.
		Preload("UserModel").
		Preload("SubCommentList").
		Where("id = ?", id).
		First(model).Error

	if err != nil {
		// 数据不存在，直接返回 nil
		return nil
	}

	res := &CommentResponse{
		ID:           model.ID,
		Content:      model.Content,
		UserID:       model.UserID,
		CreatedAt:    model.CreatedAt,
		UserNickname: model.UserModel.Nickname,
		UserAvatar:   model.UserModel.Avatar,
		ArticleID:    model.ArticleID,
		ParentID:     model.ParentID,
		DiggCount:    model.DiggCount,
		ApplyCount:   0,
		SubComments:  make([]*CommentResponse, 0),
	}

	for _, v := range model.SubCommentList {
		child := GetCommentTreeV3(v.ID)
		if child != nil {
			res.SubComments = append(res.SubComments, child)
		}
	}
	return res
}

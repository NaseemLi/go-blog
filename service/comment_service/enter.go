package commentservice

import (
	"goblog/global"
	"goblog/models"
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

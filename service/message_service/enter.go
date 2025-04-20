package messageservice

import (
	"goblog/global"
	"goblog/models"

	messagetypeenum "goblog/models/enum/message_type_enum"

	"github.com/sirupsen/logrus"
)

// 插入一条评论消息
func InsertCommentMessage(model models.CommentModel) {
	global.DB.Preload("UserModel").Preload("ArticleModel").Take(&model)
	err := global.DB.Create(&models.MessageModel{
		Type:               messagetypeenum.CommentType,
		RevUserID:          model.UserID,
		ActionUserID:       model.UserModel.ID,
		ActionUserNickname: model.UserModel.Nickname,
		ActionUserAvatar:   model.UserModel.Avatar,
		Content:            model.Content,
		ArticleID:          model.ArticleID,
		ArticleTitle:       model.ArticleModel.Title,
		CommentID:          model.ID,
	}).Error
	if err != nil {
		logrus.Error(err)
		return
	}
}

func InsertApplyMessage(model models.CommentModel) {
	//todo:回复评论的人和评论的人是同一个人时
	global.DB.Preload("ParentModel").Preload("UserModel").Preload("ArticleModel").Take(&model)
	err := global.DB.Create(&models.MessageModel{
		Type:               messagetypeenum.ApplyType,
		RevUserID:          model.ParentModel.ID,
		ActionUserID:       model.UserID,
		ActionUserNickname: model.UserModel.Nickname,
		ActionUserAvatar:   model.UserModel.Avatar,
		Content:            model.Content,
		ArticleID:          model.ArticleID,
		ArticleTitle:       model.ArticleModel.Title,
		CommentID:          model.ID,
	}).Error
	if err != nil {
		logrus.Error(err)
		return
	}
}

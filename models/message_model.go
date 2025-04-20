package models

import messagetypeenum "goblog/models/enum/message_type_enum"

type MessageModel struct {
	Model
	Type               messagetypeenum.Type `json:"type"`
	RevUserID          uint                 `json:"revUserID"`          //接收人人的 ID
	ActionUserID       uint                 `json:"ActionUserID"`       //操作人的 ID
	ActionUserNickname string               `json:"ActionUserNickname"` //操作人的昵称
	ActionUserAvatar   string               `json:"ActionUserAvatar"`   //操作人的头像
	Title              string               `json:"title"`              //标题
	Content            string               `json:"content"`            //内容
	ArticleID          uint                 `json:"articleID"`          //文章 ID
	ArticleTitle       string               `json:"articleTitle"`       //文章标题
	CommentID          uint                 `json:"commentID"`          //评论 ID
	Linktitle          string               `json:"linktitle"`          //链接标题
	LinkHref           string               `json:"linkhref"`           //链接地址
}

package sitemsgapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	messagetypeenum "goblog/models/enum/message_type_enum"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type UserMsgResponse struct {
	CommentMsgCount int `json:"commentMsgCount"`
	DiggMsgCount    int `json:"diggMsgCount"`
	PrivateMsgCount int `json:"privateMsgCount"`
	SystemMsgCount  int `json:"systemMsgCount"`
}

func (SiteMsgApi) UserMsgView(c *gin.Context) {
	claims := jwts.GetClaims(c)
	var msgList []models.MessageModel
	global.DB.Find(&msgList, "rev_user_id = ? and is_read = ?", claims.UserID, false)

	var data UserMsgResponse
	for _, model := range msgList {
		switch model.Type {
		case messagetypeenum.CommentType, messagetypeenum.ApplyType:
			data.CommentMsgCount++
		case messagetypeenum.DiggArticleType, messagetypeenum.DiggCommentType, messagetypeenum.CollectArticleType:
			data.DiggMsgCount++
		case messagetypeenum.SystermType:
			data.SystemMsgCount++
		}
	}
	//todo : 算未读私信总数 未读全局消息总数
	//过滤掉删除的,只取未读的
	var userReadMsgIDList []uint
	global.DB.Model(&models.UserGlobalNotificationModel{}).Where("user_id = ? and (is_read = ? or is_delete != ?)", claims.UserID, true, true).Select("id").Scan(&userReadMsgIDList)
	//算未读的全局消息
	var systemMsg []models.GlobalNotificationModel
	query := global.DB.Model(&models.GlobalNotificationModel{})
	if len(userReadMsgIDList) > 0 {
		query = query.Where("id NOT IN ?", userReadMsgIDList)
	}
	query.Find(&systemMsg)
	data.SystemMsgCount += len(systemMsg)

	res.OkWithData(data, c)
}

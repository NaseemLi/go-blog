package chatapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

func (ChatApi) UserChatReadView(c *gin.Context) {
	cr := middleware.GetBind[models.IDRequest](c)
	claims := jwts.GetClaims(c)
	var chat models.ChatModel
	err := global.DB.Take(&chat, cr.ID).Error
	if err != nil {
		res.FailWithMsg("消息不存在", c)
		return
	}

	var chatAc models.UserChatActionModel
	err = global.DB.Take(&chatAc, "user_id = ? and chat_id = ?", claims.UserID, chat.ID).Error
	if err != nil {
		global.DB.Create(&models.UserChatActionModel{
			UserID: claims.UserID,
			ChatID: cr.ID,
			IsRead: true,
		})
		res.OkWithMsg("消息已标记为已读", c)
		return
	}

	if chatAc.IsDelete {
		res.FailWithMsg("消息已被删除", c)
		return
	}

	global.DB.Model(&chatAc).Updates(map[string]any{
		"is_read": true,
	})

	res.OkWithMsg("消息已标记为已读", c)
}

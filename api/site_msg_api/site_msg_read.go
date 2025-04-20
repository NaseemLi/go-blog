package sitemsgapi

import (
	"fmt"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	messagetypeenum "goblog/models/enum/message_type_enum"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type SiteMsgReadRequest struct {
	ID uint `json:"id"`
	T  int8 `json:"t"`
}

func (SiteMsgApi) SiteMsgReadView(c *gin.Context) {
	cr := middleware.GetBind[SiteMsgReadRequest](c)

	claims := jwts.GetClaims(c)
	if cr.ID != 0 {
		var msg models.MessageModel
		err := global.DB.Take(&msg, "id = ? and rev_user_id = ?", cr.ID, claims.UserID).Error
		if err != nil {
			res.FailWithMsg("消息不存在", c)
			return
		}

		if msg.IsRead {
			res.FailWithMsg("消息已读", c)
			return
		}

		global.DB.Model(&msg).Where("id = ?", cr.ID).Update("is_read", true)
		res.OkWithMsg("消息读取成功", c)
		return
	}
	var typeList []messagetypeenum.Type
	switch cr.T {
	case 1:
		typeList = append(typeList, messagetypeenum.CommentType, messagetypeenum.ApplyType)
	case 2:
		typeList = append(typeList, messagetypeenum.DiggArticleType, messagetypeenum.DiggCommentType, messagetypeenum.CollectArticleType)
	case 3:
		typeList = append(typeList, messagetypeenum.SystermType)
	}

	var msgList []models.MessageModel
	global.DB.Find(&msgList, "rev_user_id = ? and type in ? and is_read = ?", claims.UserID, typeList, false)

	if len(msgList) > 0 {
		global.DB.Model(&msgList).Update("is_read", true)
	}

	res.OkWithMsg(fmt.Sprintf("批量读取 %d 条消息", len(msgList)), c)
}

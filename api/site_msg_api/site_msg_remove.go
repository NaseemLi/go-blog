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

type SiteMsgRemoveRequest struct {
	ID uint `json:"id"`
	T  int8 `json:"t"`
}

func (SiteMsgApi) SiteMsgRemoveView(c *gin.Context) {
	cr := middleware.GetBind[SiteMsgRemoveRequest](c)

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

		global.DB.Delete(&msg)
		res.OkWithMsg("消息删除成功", c)
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
	global.DB.Find(&msgList, "rev_user_id = ? and type in ?", claims.UserID, typeList)

	if len(msgList) > 0 {
		global.DB.Delete(&msgList)
	}

	res.OkWithMsg(fmt.Sprintf("批量删除 %d 条消息", len(msgList)), c)
}

package chatapi

import (
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type ChatApi struct {
}

type ChatRecordRequest struct {
	common.PageInfo
	SendUserID uint `form:"sendUserID"`
	RevUserID  uint `form:"revUserID" binding:"required"`
	Type       int8 `form:"type" binding:"required,oneof=1 2"`
}

type ChatRecordResponse struct {
	models.ChatModel
	SendUserNickname string `json:"sendUserNickname"`
	SendUserAvatar   string `json:"sendUserAvatar"`
	RevUserNickname  string `json:"revUserNickname"`
	RevUserAvatar    string `json:"revUserAvatar"`
	IsMe             bool   `json:"isMe"`
	IsRead           bool   `json:"isRead"`
}

func (ChatApi) ChatRecordView(c *gin.Context) {
	cr := middleware.GetBind[ChatRecordRequest](c)
	claims := jwts.GetClaims(c)
	var deleteIDList []uint
	var UserChatActionList []models.UserChatActionModel
	var UserChatReadMap = map[uint]bool{}
	global.DB.Find(&UserChatActionList, "user_id = ? and (is_delete = ? or is_delete is null)", cr.RevUserID, false)
	for _, v := range UserChatActionList {
		UserChatReadMap[v.ChatID] = true
	}

	switch cr.Type {
	case 1:
		//前台用户调用
		//找我删除的消息
		global.DB.Model(&models.UserChatActionModel{}).Where("user_id = ? and is_delete = ?", claims.UserID, true).Select("chat_id").Scan(&deleteIDList)

		cr.SendUserID = claims.UserID
	case 2:
		if claims.Role != enum.AdminRole {
			res.FailWithMsg("权限错误", c)
			return
		}
		if cr.SendUserID == 0 {
			res.FailWithMsg("发送者不能为空", c)
			return
		}
	}

	query := global.DB.Where("(send_user_id = ? and rev_user_id = ?) or (send_user_id = ? and rev_user_id = ?)",
		cr.SendUserID, cr.RevUserID, cr.RevUserID, cr.SendUserID)

	if len(deleteIDList) > 0 {
		query = query.Where("id not in ?", deleteIDList)
	}

	cr.Order = "created_at desc"
	_list, _, _ := common.ListQuery(models.ChatModel{}, common.Options{
		PageInfo: cr.PageInfo,
		Preloads: []string{"SendUserModel", "RevUserModel"},
		Where:    query,
	})

	var list = make([]ChatRecordResponse, 0)
	for _, model := range _list {
		item := ChatRecordResponse{
			ChatModel:        model,
			SendUserNickname: model.SendUserModel.Nickname,
			SendUserAvatar:   model.SendUserModel.Avatar,
			RevUserNickname:  model.RevUserModel.Nickname,
			RevUserAvatar:    model.RevUserModel.Avatar,
			IsRead:           UserChatReadMap[model.ID],
		}
		if model.SendUserModel.ID == claims.UserID {
			item.IsMe = true
		}
		list = append(list, item)
	}

	res.OkWithList(list, len(list), c)
}

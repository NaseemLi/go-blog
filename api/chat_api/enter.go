package chatapi

import (
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type ChatApi struct {
}

type ChatRecordRequest struct {
	common.PageInfo
	UserID uint `form:"userID" binding:"required"`
}

type ChatRecordResponse struct {
	models.ChatModel
	SendUserNickname string `json:"sendUserNickname"`
	SendUserAvatar   string `json:"sendUserAvatar"`
	RevUserNickname  string `json:"revUserNickname"`
	RevUserAvatar    string `json:"revUserAvatar"`
	IsMe             bool   `json:"isMe"`
}

func (ChatApi) ChatRecordView(c *gin.Context) {
	cr := middleware.GetBind[ChatRecordRequest](c)
	claims := jwts.GetClaims(c)
	query := global.DB.Where("(send_user_id = ? and rev_user_id = ?) or (send_user_id = ? and rev_user_id = ?)", cr.UserID, claims.UserID, claims.UserID, cr.UserID)

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
		}
		if model.SendUserModel.ID == claims.UserID {
			item.IsMe = true
		}
		list = append(list, item)
	}

	res.OkWithList(list, len(list), c)
}

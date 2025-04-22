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

func (ChatApi) UserChatDeleteView(c *gin.Context) {
	cr := middleware.GetBind[models.RemoveRequest](c)
	claims := jwts.GetClaims(c)

	var chatList []models.ChatModel
	global.DB.Find(&chatList, "id in ?", cr.IDList)

	chatMap := common.ScanMapV2(models.UserChatActionModel{}, common.ScanOption{
		Where: global.DB.Where("user_id = ? and chat_id in ?", claims.UserID, cr.IDList),
		Key:   "ChatID",
	})

	var addChatAc []models.UserChatActionModel
	var updateChatAcIdList []uint
	for _, v := range chatList {
		//判断是否删过
		chat, ok := chatMap[v.ID]
		if !ok {
			//说明没有删除过
			addChatAc = append(addChatAc, models.UserChatActionModel{
				UserID:   claims.UserID,
				ChatID:   v.ID,
				IsDelete: true,
			})
			continue
		}
		if chat.IsDelete {
			continue
		}
		updateChatAcIdList = append(updateChatAcIdList, chat.ID)
	}
	if len(addChatAc) > 0 {
		err := global.DB.Create(&addChatAc).Error
		if err != nil {
			res.FailWithMsg("消息删除失败", c)
			return
		}
	}
	if len(updateChatAcIdList) > 0 {
		err := global.DB.Model(&models.UserChatActionModel{}).
			Where("id in ?", updateChatAcIdList).
			Updates(map[string]any{
				"is_delete": true,
			}).Error
		if err != nil {
			res.FailWithMsg("消息删除失败", c)
			return
		}
	}

	var DeleteList []uint
	for _, v := range addChatAc {
		DeleteList = append(DeleteList, v.ChatID)
	}
	DeleteList = append(DeleteList, updateChatAcIdList...)

	res.OkWithList(DeleteList, len(DeleteList), c)
}

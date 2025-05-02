package sitemsgapi

import (
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	messagetypeenum "goblog/models/enum/message_type_enum"
	relationshipenum "goblog/models/enum/relationship_enum"
	focusservice "goblog/service/focus_service"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type SiteMsgListRequest struct {
	common.PageInfo
	T int8 `form:"t" binding:"required,oneof=1 2 3"` //1.评论和回复 2.赞和收藏 3.系统消息
}
type SiteMsgListResponse struct {
	models.MessageModel
	Relation relationshipenum.Relation `json:"relation"`
}

func (SiteMsgApi) SiteMsgListView(c *gin.Context) {
	cr := middleware.GetBind[SiteMsgListRequest](c)

	var typeList []messagetypeenum.Type
	switch cr.T {
	case 1:
		typeList = append(typeList, messagetypeenum.CommentType, messagetypeenum.ApplyType)
	case 2:
		typeList = append(typeList, messagetypeenum.DiggArticleType, messagetypeenum.DiggCommentType, messagetypeenum.CollectArticleType)
	case 3:
		typeList = append(typeList, messagetypeenum.SystermType)
	}

	claims := jwts.GetClaims(c)
	_list, _, _ := common.ListQuery(models.MessageModel{
		RevUserID: claims.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Where:    global.DB.Where("type in ?", typeList),
	})

	var userIDList []uint
	for _, v := range _list {
		if v.ActionUserID != 0 {
			userIDList = append(userIDList, v.ActionUserID)
		}
	}

	var m = make(map[uint]relationshipenum.Relation)
	if len(userIDList) > 0 {
		m = focusservice.CalcUserPatchRelationship(claims.UserID, userIDList)
	}

	var list = make([]SiteMsgListResponse, 0)
	for _, model := range _list {
		list = append(list, SiteMsgListResponse{
			MessageModel: model,
			Relation:     m[model.ActionUserID],
		})
	}

	res.OkWithList(list, len(list), c)
}

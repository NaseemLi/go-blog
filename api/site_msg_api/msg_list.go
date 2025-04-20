package sitemsgapi

import (
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	messagetypeenum "goblog/models/enum/message_type_enum"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type SiteMsgListRequest struct {
	common.PageInfo
	T int8 `form:"t" binding:"required,oneof=1 2 3"` //1.评论和回复 2.赞和收藏 3.系统消息
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
	list, _, _ := common.ListQuery(models.MessageModel{
		RevUserID: claims.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Where:    global.DB.Where("type in ?", typeList),
	})
	res.OkWithList(list, len(list), c)
}

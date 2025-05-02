package articleapi

import (
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	relationshipenum "goblog/models/enum/relationship_enum"
	focusservice "goblog/service/focus_service"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type AuthRecommentResponse struct {
	UserID       uint   `json:"userID"`
	UserNickname string `json:"userNickname"`
	UserAvatar   string `json:"userAvatar"`
	UserAbstract string `json:"userAbstract"`
}

func (ArticleApi) AuthRecommentView(c *gin.Context) {
	cr := middleware.GetBind[common.PageInfo](c)

	var count int
	var userIDList []uint
	global.DB.Model(models.ArticleModel{}).Group("user_id").
		Select("COUNT(*)").Scan(&count)
	global.DB.Model(models.ArticleModel{}).Group("user_id").
		Offset(cr.GetOffset()).Limit(cr.GetLimit()).
		Select("user_id").Scan(&userIDList)

	claims, err := jwts.ParseTokenByGin(c)
	if err == nil && claims != nil {
		m := focusservice.CalcUserPatchRelationship(claims.UserID, userIDList)
		userIDList = []uint{}
		for u, relationn := range m {
			if relationn == relationshipenum.RelationStranger || relationn == relationshipenum.RelationFans {
				userIDList = append(userIDList, u)
			}
		}
	}
	var userList []models.UserModel
	global.DB.Find(&userList, "id in ?", userIDList)

	var list = make([]AuthRecommentResponse, 0)
	for _, v := range userList {
		list = append(list, AuthRecommentResponse{
			UserID:       v.ID,
			UserNickname: v.Nickname,
			UserAvatar:   v.Avatar,
			UserAbstract: v.Abstract,
		})
	}
	res.OkWithList(list, count, c)
}

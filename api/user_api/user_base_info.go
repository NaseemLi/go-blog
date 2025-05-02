package userapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	relationshipenum "goblog/models/enum/relationship_enum"
	focusservice "goblog/service/focus_service"
	redisuser "goblog/service/redis_service/redis_user"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type UserBaseInfoResponse struct {
	UserID       uint                      `json:"userID"`
	CodeAge      int                       `json:"codeAge"`
	Avatar       string                    `json:"avatar"`
	Nickname     string                    `json:"nickname"`
	LookCount    int                       `json:"lookCount"`
	ArticleCount int                       `json:"articleCount"`
	FansCount    int                       `json:"fansCount"`
	FollowCount  int                       `json:"followCount"`
	Place        string                    `json:"place"`
	OpenFollow   bool                      `json:"openFollow"`  // 是否公开关注
	OpenCollect  bool                      `json:"openCollect"` // 是否公开收藏
	OpenFans     bool                      `json:"openFans"`    // 是否公开粉丝
	HomeStyleID  uint                      `json:"homeStyleID"` // 主页风格
	Relation     relationshipenum.Relation `json:"relation"`    // 关系
}

func (UserApi) UserBaseInfoView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var user models.UserModel
	err = global.DB.Preload("UserConfModel").Preload("ArticleList").Take(&user, cr.ID).Error
	if err != nil {
		res.FailWithMsg("不存在的用户", c)
		return
	}

	data := UserBaseInfoResponse{
		UserID:       user.ID,
		CodeAge:      user.GetCodeAge(),
		Avatar:       user.Avatar,
		Nickname:     user.Nickname,
		LookCount:    user.UserConfModel.LookCount + redisuser.GetCacheLook(cr.ID),
		ArticleCount: len(user.ArticleList),
		FansCount:    0,
		FollowCount:  0,
		Place:        user.Addr,
		OpenFollow:   user.UserConfModel.OpenFollow,
		OpenCollect:  user.UserConfModel.OpenCollect,
		OpenFans:     user.UserConfModel.OpenFans,
		HomeStyleID:  user.UserConfModel.HomeStyleID,
	}

	claims, err := jwts.ParseTokenByGin(c)
	if err == nil && claims != nil {
		data.Relation = focusservice.CalcUserRelationship(claims.UserID, cr.ID)
	}

	var focusList []models.UserFocusModel
	global.DB.Find(&focusList, "user_id = ? OR focus_user_id = ?", cr.ID, cr.ID)
	for _, model := range focusList {
		if model.UserID == cr.ID {
			data.FollowCount++
		}
		if model.FocusUserID == cr.ID {
			data.FansCount++
		}

	}
	redisuser.SetCacheLook(cr.ID, true)

	res.OkWithData(data, c)
}

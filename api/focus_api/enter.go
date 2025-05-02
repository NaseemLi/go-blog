package focusapi

import (
	"fmt"
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	relationshipenum "goblog/models/enum/relationship_enum"
	focusservice "goblog/service/focus_service"
	"goblog/utils/jwts"
	"time"

	"github.com/gin-gonic/gin"
)

type FocusApi struct {
}

type FocusUserRequest struct {
	FocusUserID uint `json:"focusUserID" binding:"required"`
}

func (FocusApi) FocusUserView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserRequest](c)
	claims := jwts.GetClaims(c)

	//自己不能关注自己
	if claims.UserID == cr.FocusUserID {
		res.FailWithMsg("你正在关注自己", c)
		return
	}

	//查关注的用户是否存在
	var user models.UserModel
	err := global.DB.Take(&user, cr.FocusUserID).Error
	if err != nil {
		res.FailWithMsg("关注用户不存在", c)
		return
	}

	//查之前是否已经关注
	var userFocus models.UserFocusModel
	err = global.DB.Take(&userFocus, "user_id = ? and focus_user_id = ?", claims.UserID, cr.FocusUserID).Error
	if err == nil {
		// 查到了记录，说明已经关注
		res.FailWithMsg("请勿重复关注", c)
		return
	}

	//添加关注
	global.DB.Create(&models.UserFocusModel{
		UserID:      claims.UserID,
		FocusUserID: cr.FocusUserID,
	})

	res.OkWithMsg("关注成功", c)
}

type FocusUserListRequest struct {
	common.PageInfo
	FocusUserID uint `form:"focusUserID"`
	UserID      uint `form:"userID"`
}

type UserListResponse struct {
	UserID       uint                      `json:"userID"`
	UserNickname string                    `json:"userNickname"`
	UserAvatar   string                    `json:"userAvatar"`
	UserAbstract string                    `json:"userAbstract"`
	Relationship relationshipenum.Relation `json:"relationship"`
	CreateAt     time.Time                 `json:"createAt"`
}

// 我的关注和用户的关注
func (FocusApi) FocusUserListView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserListRequest](c)
	claims, err := jwts.ParseTokenByGin(c)

	if cr.UserID != 0 {
		//如果传了用户id,就查这个人关注的用户列表
		var userconf models.UserConfModel
		err1 := global.DB.Take(&userconf, "user_id = ?", cr.UserID).Error
		if err1 != nil {
			res.FailWithMsg("用户配置信息不存在", c)
			return
		}
		if !userconf.OpenFollow {
			res.FailWithMsg("用户没有公开我的关注", c)
			return
		}

		//如果没登录,只允许查询第一页
		if err != nil || claims == nil {
			if cr.Limit > 10 || cr.Page > 1 {
				res.FailWithMsg("登录后查看更多", c)
				return
			}
		}
	} else {
		if err != nil || claims == nil {
			res.FailWithMsg("无效的token,请登录", c)
			return
		}
		cr.UserID = claims.UserID
	}

	query := global.DB.Where("")
	if cr.Key != "" {
		//模糊匹配用户
		var userIDList []uint
		global.DB.Model(&models.UserModel{}).Where("nickname like ?", fmt.Sprintf("%%%s%%", cr.Key)).Select("id").Scan(&userIDList)
		if len(userIDList) > 0 {
			query.Where("focus_user_id in ?", userIDList)
		}
	}

	_list, _, _ := common.ListQuery(models.UserFocusModel{
		FocusUserID: cr.FocusUserID,
		UserID:      cr.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Preloads: []string{"FocusUserModel"},
		Where:    query,
	})

	var m = map[uint]relationshipenum.Relation{}
	if err == nil && claims != nil {
		var userIDList []uint
		for _, v := range _list {
			userIDList = append(userIDList, v.FocusUserID)
		}
		m = focusservice.CalcUserPatchRelationship(claims.UserID, userIDList)
	}

	var list = make([]UserListResponse, 0)
	for _, model := range _list {
		list = append(list, UserListResponse{
			UserID:       model.FocusUserModel.ID,
			UserNickname: model.FocusUserModel.Nickname,
			UserAvatar:   model.FocusUserModel.Avatar,
			UserAbstract: model.FocusUserModel.Abstract,
			Relationship: m[model.FocusUserID],
			CreateAt:     model.CreatedAt,
		})
	}
	res.OkWithList(list, len(list), c)
}

// 我的粉丝和用户的粉丝
func (FocusApi) FansUserListView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserListRequest](c)
	claims, err := jwts.ParseTokenByGin(c)

	if cr.UserID != 0 {
		//如果传了用户id,就查这个人的粉丝列表
		var userconf models.UserConfModel
		err1 := global.DB.Take(&userconf, "user_id = ?", cr.UserID).Error
		if err1 != nil {
			res.FailWithMsg("用户配置信息不存在", c)
			return
		}
		if !userconf.OpenFans {
			res.FailWithMsg("用户没有公开我的粉丝", c)
			return
		}

		if err != nil || claims == nil {
			//如果没登录,只允许查询第一页
			if cr.Limit > 10 || cr.Page > 1 {
				res.FailWithMsg("登录后查看更多", c)
				return
			}
		}
	} else {
		if err != nil || claims == nil {
			res.FailWithMsg("无效的token,请登录", c)
			return
		}
		cr.UserID = claims.UserID
	}

	query := global.DB.Where("")
	if cr.Key != "" {
		//模糊匹配用户
		var userIDList []uint
		global.DB.Model(&models.UserModel{}).Where("nickname like ?", fmt.Sprintf("%%%s%%", cr.Key)).Select("id").Scan(&userIDList)
		if len(userIDList) > 0 {
			query.Where("user_id in ?", userIDList)
		}
	}

	_list, _, _ := common.ListQuery(models.UserFocusModel{
		FocusUserID: cr.UserID,
		UserID:      cr.FocusUserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Preloads: []string{"FocusUserModel"},
		Where:    query,
	})

	var m = map[uint]relationshipenum.Relation{}
	if err == nil && claims != nil {
		var userIDList []uint
		for _, v := range _list {
			userIDList = append(userIDList, v.UserID)
		}
		m = focusservice.CalcUserPatchRelationship(claims.UserID, userIDList)
	}

	var list = make([]UserListResponse, 0)
	for _, model := range _list {
		list = append(list, UserListResponse{
			UserID:       model.UserID,
			UserNickname: model.FocusUserModel.Nickname,
			UserAvatar:   model.FocusUserModel.Avatar,
			UserAbstract: model.FocusUserModel.Abstract,
			Relationship: m[model.FocusUserID],
			CreateAt:     model.CreatedAt,
		})
	}
	res.OkWithList(list, len(list), c)
}

type UnFocusUserRequest struct {
	FocusUserID uint `json:"focusUserID" binding:"required"`
}

func (FocusApi) UnFocusUserView(c *gin.Context) {
	cr := middleware.GetBind[UnFocusUserRequest](c)
	claims := jwts.GetClaims(c)

	//自己不能取关自己
	if claims.UserID == cr.FocusUserID {
		res.FailWithMsg("失败!你正在取关自己", c)
		return
	}

	//查取关的用户是否存在
	var user models.UserModel
	err := global.DB.Take(&user, cr.FocusUserID).Error
	if err != nil {
		res.FailWithMsg("取关用户不存在", c)
		return
	}

	//查之前是否已经关注
	var userFocus models.UserFocusModel
	err = global.DB.Take(&userFocus, "user_id = ? and focus_user_id = ?", claims.UserID, cr.FocusUserID).Error
	if err != nil {
		// 查到了记录，说明已经关注
		res.FailWithMsg("未关注此用户", c)
		return
	}

	//取消关注
	global.DB.Delete(&userFocus)

	res.OkWithMsg("取关成功", c)
}

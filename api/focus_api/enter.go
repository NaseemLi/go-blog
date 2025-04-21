package focusapi

import (
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
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

type FocusUserListResponse struct {
	FocusUserID       uint      `json:"focusUserID"`
	FocusUserNickname string    `json:"focusUserNickname"`
	FocusUserAvatar   string    `json:"focusUserAvatar"`
	FocusUserAbstract string    `json:"focusUserAbstract"`
	CreateAt          time.Time `json:"createAt"`
}

// 我的关注和用户的关注
func (FocusApi) FocusUserListView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserListRequest](c)
	if cr.UserID != 0 {
		//如果传了用户id,就查这个人关注的用户列表
		var userconf models.UserConfModel
		err := global.DB.Take(&userconf, "user_id = ?", cr.UserID).Error
		if err != nil {
			res.FailWithMsg("用户配置信息不存在", c)
			return
		}
		if !userconf.OpenFollow {
			res.FailWithMsg("用户没有公开我的关注", c)
			return
		}
	} else {
		claims, err := jwts.ParseTokenByGin(c)
		if err != nil || claims == nil {
			res.FailWithMsg("无效的token,请登录", c)
			return
		}
		cr.UserID = claims.UserID
	}

	_list, _, _ := common.ListQuery(models.UserFocusModel{
		FocusUserID: cr.FocusUserID,
		UserID:      cr.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Preloads: []string{"FocusUserModel"},
	})

	var list = make([]FocusUserListResponse, 0)
	for _, model := range _list {
		list = append(list, FocusUserListResponse{
			FocusUserID:       model.FocusUserModel.ID,
			FocusUserNickname: model.FocusUserModel.Nickname,
			FocusUserAvatar:   model.FocusUserModel.Avatar,
			FocusUserAbstract: model.FocusUserModel.Abstract,
			CreateAt:          model.CreatedAt,
		})
	}
	res.OkWithList(list, len(list), c)
}

type FansUserListResponse struct {
	FansUserID       uint      `json:"fansUserID"`
	FansUserNickname string    `json:"fansUserNickname"`
	FansUserAvatar   string    `json:"fansUserAvatar"`
	FansUserAbstract string    `json:"fansUserAbstract"`
	CreateAt         time.Time `json:"createAt"`
}

// 我的粉丝和用户的粉丝
func (FocusApi) FansUserListView(c *gin.Context) {
	cr := middleware.GetBind[FocusUserListRequest](c)
	if cr.UserID != 0 {
		//如果传了用户id,就查这个人的粉丝列表
		var userconf models.UserConfModel
		err := global.DB.Take(&userconf, "user_id = ?", cr.UserID).Error
		if err != nil {
			res.FailWithMsg("用户配置信息不存在", c)
			return
		}
		if !userconf.OpenFans {
			res.FailWithMsg("用户没有公开我的粉丝", c)
			return
		}
	} else {
		claims, err := jwts.ParseTokenByGin(c)
		if err != nil || claims == nil {
			res.FailWithMsg("无效的token,请登录", c)
			return
		}
		cr.UserID = claims.UserID
	}

	_list, _, _ := common.ListQuery(models.UserFocusModel{
		FocusUserID: cr.UserID,
		UserID:      cr.FocusUserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Preloads: []string{"FocusUserModel"},
	})

	var list = make([]FansUserListResponse, 0)
	for _, model := range _list {
		list = append(list, FansUserListResponse{
			FansUserID:       model.UserID,
			FansUserNickname: model.FocusUserModel.Nickname,
			FansUserAvatar:   model.FocusUserModel.Avatar,
			FansUserAbstract: model.FocusUserModel.Abstract,
			CreateAt:         model.CreatedAt,
		})
	}
	res.OkWithList(list, len(list), c)
}

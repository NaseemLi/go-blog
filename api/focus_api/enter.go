package focusapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type FocusApi struct {
}

type FocusUserRequest struct {
	FocusUserID uint `json:"focusUserID" binding:"required"`
}

func (FocusApi) FocusUserApi(c *gin.Context) {
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
	if err != nil {
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

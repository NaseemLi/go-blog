package userapi

import (
	"fmt"
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	"goblog/models/enum"
	qqservice "goblog/service/qq_service"
	userservice "goblog/service/user_service"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type QQLoginRequest struct {
	Code string `json:"code" binding:"required"`
}

func (UserApi) QQLoginView(c *gin.Context) {
	var cr QQLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	info, err := qqservice.GetUserInfo(cr.Code)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, "open_id = ?", info.OpenID).Error
	if err != nil {
		//创建用户
		uname := base64Captcha.RandText(5, "1234567890")
		user = models.UserModel{
			Username:       fmt.Sprintf("b_%s", uname),
			Nickname:       info.Nickname,
			Avatar:         info.Avatar,
			RegisterSource: enum.RegisterQQSourceType,
			OpenID:         info.OpenID,
			Role:           enum.UserRole,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			logrus.Error(err)
			res.FailWithMsg("qq 登录失败", c)
			return
		}
	}
	//颁发 token
	token, _ := jwts.GetToken(jwts.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})

	userservice.NewUserService(user).UserLogin(c)
	res.OkWithData(token, c)
}

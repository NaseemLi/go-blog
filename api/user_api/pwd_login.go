package userapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	userservice "goblog/service/user_service"
	"goblog/utils/jwts"
	"goblog/utils/pwd"

	"github.com/gin-gonic/gin"
)

type PwdLoginrequest struct {
	Val      string `json:"val"`
	Password string `json:"password"`
}

func (UserApi) PwdLoginApi(c *gin.Context) {
	var cr PwdLoginrequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	if !global.Config.Site.Login.UsernamePwdLogin {
		res.FailWithMsg("站点未启用密码登录", c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, "(username = ? or email = ?) and password <> ''", cr.Val, cr.Val).Error
	if err != nil {
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	if !pwd.CompareHashAndPassword(user.Password, cr.Password) {
		res.FailWithMsg("用户名或密码错误", c)
		return
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

package userapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/models/enum"
	"goblog/utils/jwts"
	"goblog/utils/pwd"

	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"oldPwd"`
	Pwd    string `json:"pwd"`
}

func (UserApi) UpdatePasswordView(c *gin.Context) {
	var cr UpdatePasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	claims := jwts.GetClaims(c)
	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	//邮箱注册的,绑了邮箱的
	if !(user.RegisterSource == enum.RegisterEmailSourceType || user.Email != "") {
		res.FailWithMsg("仅支持邮箱注册或绑定邮箱的用户修改密码", c)
		return
	}

	//校验之前的密码
	if !pwd.CompareHashAndPassword(user.Password, cr.OldPwd) {
		res.FailWithMsg("旧密码错误", c)
		return
	}

	hashPwd, _ := pwd.GenerateFromPassword(cr.Pwd)
	global.DB.Model(&user).Update("password", hashPwd)
	res.OkWithMsg("密码修改成功", c)
}

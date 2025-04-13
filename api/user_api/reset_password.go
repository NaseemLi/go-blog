package userapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	"goblog/models/enum"
	emailstore "goblog/utils/email_store"
	"goblog/utils/pwd"

	"github.com/gin-gonic/gin"
)

type ResetPasswordRequest struct {
	EmailID   string `json:"emailID" binding:"required"`
	EmailCode string `json:"emailCode" binding:"required"`
	Pwd       string `json:"pwd" binding:"required"`
}

func (UserApi) ResetPasswordView(c *gin.Context) {
	var cr ResetPasswordRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, c)
		return
	}

	// 从 EmailID 拿回原始 email/code
	v, ok := global.EmailVerifyStore.Load(cr.EmailID)
	if !ok {
		res.FailWithMsg("验证码已过期或无效", c)
		return
	}

	info := v.(emailstore.EmailStoreInfo)

	// 校验验证码
	if info.Code != cr.EmailCode {
		res.FailWithMsg("验证码错误", c)
		return
	}

	// 查用户
	var user models.UserModel
	err := global.DB.Take(&user, "email = ?", info.Email).Error
	if err != nil {
		res.FailWithMsg("该邮箱未注册", c)
		return
	}

	if user.RegisterSource != enum.RegisterEmailSourceType {
		res.FailWithMsg("非邮箱注册用户不能重置密码", c)
		return
	}

	// 修改密码
	hashPwd, _ := pwd.GenerateFromPassword(cr.Pwd)
	global.DB.Model(&user).Update("password", hashPwd)

	// 删除验证码缓存
	global.EmailVerifyStore.Delete(cr.EmailID)

	res.OkWithMsg("密码重置成功", c)
}

package userapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	emailservice "goblog/service/email_service"
	emailstore "goblog/utils/email_store"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/sirupsen/logrus"
)

type SendEmailRequest struct {
	Type  int8   `json:"type" binding:"oneof=1 2"` //1.注册 2.重置密码
	Email string `json:"email" binding:"required"`
}

type SendEmailResponse struct {
	EmailID string `json:"emailID"`
}

func (UserApi) SendEmailView(c *gin.Context) {
	var cr SendEmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}
	code := base64Captcha.RandText(4, "1234567890")
	id := base64Captcha.RandomId()
	switch cr.Type {
	case 1:
		var user models.UserModel
		err = global.DB.Take(&user, "email = ?", cr.Email).Error
		if err == nil {
			res.FailWithMsg("该邮箱已使用", c)
		}
		err = emailservice.SendRegisterCode(cr.Email, code)
	case 2:
		err = emailservice.SendResetPwdCode(cr.Email, code)
	}
	if err != nil {
		logrus.Errorf("邮件发送失败 %s", err)
		res.FailWithMsg("邮件发送失败", c)
		return
	}
	global.CaptchaStore.Set(id, code)
	global.EmailVerifyStore.Store(id, emailstore.EmailStoreInfo{
		Email: cr.Email,
		Code:  code,
	})
	res.OkWithData(SendEmailResponse{
		EmailID: id,
	}, c)
}

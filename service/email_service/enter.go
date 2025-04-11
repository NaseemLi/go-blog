package emailservice

import (
	"fmt"
	"goblog/global"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

func SendRegisterCode(to string, code string) error {
	subject := "账号注册"
	text := fmt.Sprintf("你正在进行账号注册,这是你的验证码%s,10分钟内有效", code)
	return SendEmail(to, subject, text)
}

func SendResetPwdCode(to string, code string) error {
	subject := "密码重置"
	text := fmt.Sprintf("你正在进行账号密码重置,这是你的验证码%s,10分钟内有效", code)
	return SendEmail(to, subject, text)
}

func SendEmail(to, subject, text string) (err error) {
	em := global.Config.Email
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", em.SendNickname, em.SendEmail)
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(text)
	err1 := e.Send(fmt.Sprintf("%s:%d", em.Domain, em.Port), smtp.PlainAuth("", em.SendEmail, em.AuthCode, em.Domain))
	if err1 != nil && !strings.Contains(err1.Error(), "short response:") {
		return err1
	}
	return nil
}

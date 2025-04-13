package userapi

import (
	"goblog/common/res"
	"goblog/global"
	emailstore "goblog/utils/email_store"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type BindEmailRequest struct {
	EmailID   string `json:"emailID" binding:"required"`
	EmailCode string `json:"emailCode" binding:"required"`
}

func (UserApi) BindEmailView(c *gin.Context) {
	if !global.Config.Site.Login.EmailLogin {
		res.FailWithMsg("站点未启用邮箱注册", c)
		return
	}

	var cr BindEmailRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, c)
		return
	}

	// 从 EmailVerifyStore 拿到邮箱和验证码
	v, ok := global.EmailVerifyStore.Load(cr.EmailID)
	if !ok {
		res.FailWithMsg("验证码已过期或无效", c)
		return
	}

	info := v.(emailstore.EmailStoreInfo)
	if info.Code != cr.EmailCode {
		res.FailWithMsg("验证码错误", c)
		return
	}

	// 从 token 获取当前用户
	claims := jwts.GetClaims(c)
	if claims == nil {
		res.FailWithMsg("用户未登录或 token 无效", c)
		return
	}
	user, err := claims.GetUser()
	if err != nil {
		res.FailWithMsg("不存在的用户", c)
		return
	}

	// 绑定邮箱
	global.DB.Model(&user).Update("email", info.Email)

	// 删除验证码，防止复用
	global.EmailVerifyStore.Delete(cr.EmailID)

	res.OkWithMsg("邮箱绑定成功", c)
}

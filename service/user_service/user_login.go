package userservice

import (
	"goblog/core"
	"goblog/global"
	"goblog/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (u UserService) UserLogin(c *gin.Context) {
	ip := c.ClientIP()
	addr := core.GetipADDR(ip)
	ua := c.GetHeader("User-Agent")
	err := global.DB.Create(&models.UserLoginModel{
		UserID: u.userModel.ID,
		IP:     ip,
		Addr:   addr,
		UA:     ua,
	}).Error
	if err != nil {
		logrus.Errorf("用户登录日志写入失败%s", err)
	}
}

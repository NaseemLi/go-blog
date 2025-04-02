package siteapi

import (
	"fmt"
	"goblog/models/enum"
	logservice "goblog/service/log_service"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Siteapi struct {
}

func (Siteapi) SiteInfoView(c *gin.Context) {
	fmt.Println("1")
	logservice.NewLoginSuccess(c, enum.UserPwdLoginType)
	logservice.NewLoginFail(c, enum.UserPwdLoginType, "用户不存在", "fengfeng", "1234")
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "站点信息",
	})
}

type SiteUpdateRequest struct {
	Name string `json:"name"`
}

func (Siteapi) SiteUpdateView(c *gin.Context) {
	log := logservice.NewActionLogByGin(c)
	var cr SiteUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		logrus.Errorf(err.Error())
	}

	log.Save()
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "站点信息",
	})
}

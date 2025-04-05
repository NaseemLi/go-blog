package siteapi

import (
	"fmt"
	"goblog/models/enum"
	logservice "goblog/service/log_service"
	"time"

	"github.com/gin-gonic/gin"
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
	Name string `json:"name" binding:"required"`
}

func (Siteapi) SiteUpdateView(c *gin.Context) {
	log := logservice.Getlog(c)
	log.ShowRequest()
	log.ShowRequestHeader()
	log.ShowResponseHeader()
	log.SetImage("/xxx/xxx")
	log.SetLink("yaml学习地址", "https://www.fengfengzhidao.com")
	log.ShowResponse()
	log.SetTitle("更新站点")
	log.SetItemInfo("请求时间", time.Now())

	var cr SiteUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		log.SetError("参数绑定失败", err)
	}
	log.SetItemInfo("结构体", cr)

	id := log.Save()
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "站点信息",
		"data": id,
	})
}

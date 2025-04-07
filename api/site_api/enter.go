package siteapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

type Siteapi struct {
}

type SiteInfoRequest struct {
	Name string `uri:"name"`
}

func (Siteapi) SiteInfoView(c *gin.Context) {
	var cr SiteInfoRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	if cr.Name == "site" {
		global.Config.Site.About.Version = global.Version
		res.OkWithData(global.Config.Site, c)
	}

	//判断角色是不是
	middleware.AdminMiddelware(c)

	_, ok := c.Get("claims")
	if !ok {
		return
	}

	var data any
	switch cr.Name {
	case "email":
		rep := global.Config.Email
		rep.AuthCode = "*******"
		data = rep
	case "qq":
		rep := global.Config.QQ
		rep.AppKey = "*******"
		data = rep
	case "qiniu":
		rep := global.Config.QiNiu
		rep.SecretKey = "*******"
		data = rep
	case "ai":
		rep := global.Config.Ai
		rep.SecretKey = "*******"
		data = rep
	default:
		res.FailWithMsg("不存在的配置", c)
		return
	}

	res.OkWithData(data, c)
	return
}

func (Siteapi) SiteInfoQQView(c *gin.Context) {
	res.OkWithData(global.Config.QQ.Url(), c)
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  string `json:"age" binding:"required" label:"年龄"`
}

func (Siteapi) SiteUpdateView(c *gin.Context) {
	//log := logservice.Getlog(c)
	var cr SiteUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	res.OkWithMsg("更新成功", c)
}

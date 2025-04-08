package siteapi

import (
	"goblog/common/res"
	"goblog/conf"
	"goblog/core"
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
	var cr SiteInfoRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	var rep any
	switch cr.Name {
	case "site":
		var data conf.Site
		err = c.ShouldBindJSON(&data)
		rep = data
	case "email":
		var data conf.Email
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qq":
		var data conf.QQ
		err = c.ShouldBindJSON(&data)
		rep = data
	case "qiniu":
		var data conf.QiNiu
		err = c.ShouldBindJSON(&data)
		rep = data
	case "ai":
		var data conf.Ai
		err = c.ShouldBindJSON(&data)
		rep = data
	default:
		res.FailWithMsg("不存在配置", c)
	}
	if err != nil {
		res.FailWithError(err, c)
		return
	}

	switch s := rep.(type) {
	case conf.Site:
		//判断站点信息更新前端部分
		err = UpdateSite(s)
		if err != nil {
			res.FailWithError(err, c)
			return
		}
		global.Config.Site = s
	case conf.Email:
		if s.AuthCode == "******" {
			s.AuthCode = global.Config.Email.AuthCode
		}
		global.Config.Email = s
	case conf.QQ:
		if s.AppKey == "******" {
			s.AppKey = global.Config.QQ.AppKey
		}
		global.Config.QQ = s
	case conf.QiNiu:
		if s.SecretKey == "******" {
			s.SecretKey = global.Config.QiNiu.SecretKey
		}
		global.Config.QiNiu = s
	case conf.Ai:
		if s.SecretKey == "******" {
			s.SecretKey = global.Config.Ai.SecretKey
		}
		global.Config.Ai = s
	}

	//改配置文件
	core.SetConf()

	res.OkWithMsg("更新站点配置成功", c)
	return
}

func UpdateSite(site conf.Site) error {
	return nil
}

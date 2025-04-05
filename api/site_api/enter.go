package siteapi

import (
	"goblog/common/res"

	"github.com/gin-gonic/gin"
)

type Siteapi struct {
}

func (Siteapi) SiteInfoView(c *gin.Context) {
	res.OkWithData("xx", c)
	return
}

type SiteUpdateRequest struct {
	Name string `json:"name" binding:"required"`
}

func (Siteapi) SiteUpdateView(c *gin.Context) {
	//log := logservice.Getlog(c)
	var cr SiteUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	res.OkWithMsg("更新成功", c)
}

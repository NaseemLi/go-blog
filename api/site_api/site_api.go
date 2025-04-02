package siteapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Siteapi struct {
}

func (Siteapi) SiteInfoView(c *gin.Context) {
	fmt.Println("1")
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "站点信息",
	})
}

package middleware

import (
	"goblog/common/res"
	"goblog/models/enum"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

func AuthMiddelware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
	return
}

func AdminMiddelware(c *gin.Context) {
	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.FailWithError(err, c)
		c.Abort()
		return
	}
	if claims.Role != enum.AdminRole {
		res.FailWithMsg("权限错误", c)
		c.Abort()
		return
	}
	c.Set("claims", claims)
}

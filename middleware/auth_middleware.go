package middleware

import (
	"goblog/common/res"
	"goblog/models/enum"
	redisjwt "goblog/service/redis_service/redis_jwt"
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
	blcType, ok := redisjwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(blcType.String(), c)
		return
	}
	c.Set("claims", claims)

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
	blcType, ok := redisjwt.HasTokenBlackByGin(c)
	if ok {
		res.FailWithMsg(blcType.Msg(), c)
		c.Abort()
		return
	}

	c.Set("claims", claims)
}

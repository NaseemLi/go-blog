package userapi

import (
	"goblog/common/res"
	redisjwt "goblog/service/redis_service/redis_jwt"

	"github.com/gin-gonic/gin"
)

func (UserApi) LogOutView(c *gin.Context) {
	token := c.Request.Header.Get("token")
	redisjwt.TokenBlack(token, redisjwt.UserBlackType)

	res.OkWithMsg("退出成功", c)
}

package userapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/models"
	"goblog/models/enum"
	"goblog/utils/jwts"
	"math"
	"time"

	"github.com/gin-gonic/gin"
)

type UserDetailResponse struct {
	ID             uint                    `json:"id"`
	CreatedAt      time.Time               `json:"createdAt"`
	Username       string                  `json:"username"`
	Nickname       string                  `json:"nickname"`
	Avatar         string                  `json:"avatar"`
	Abstract       string                  `json:"abstract"`
	RegisterSource enum.RegisterSourceType `json:"registerSource"` // 注册来源
	CodeAge        int                     `json:"codeAge"`        // 码龄
	models.UserConfModel
}

func (UserApi) UserDetailView(c *gin.Context) {
	claims := jwts.GetClaims(c)
	var user models.UserModel
	err := global.DB.Preload("UserConfModel").Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	sub := time.Now().Sub(user.CreatedAt)
	codeAge := int(math.Ceil(sub.Hours() / 24 / 365))

	var data = UserDetailResponse{
		ID:             user.ID,
		CreatedAt:      user.CreatedAt,
		Username:       user.Username,
		Nickname:       user.Nickname,
		Avatar:         user.Avatar,
		Abstract:       user.Abstract,
		RegisterSource: user.RegisterSource,
		CodeAge:        codeAge,
	}
	if user.UserConfModel != nil {
		data.UserConfModel = *user.UserConfModel
	}

	res.OkWithData(data, c)
}

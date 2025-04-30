package userapi

import (
	"goblog/common"
	"goblog/common/res"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	"time"

	"github.com/gin-gonic/gin"
)

type UserListRequest struct {
	common.PageInfo
}

type UserListRespone struct {
	ID            uint          `json:"id"`
	Nickname      string        `json:"nickname"`
	Username      string        `json:"username"`
	Avatar        string        `json:"avatar"`
	IP            string        `json:"ip"`
	Addr          string        `json:"addr"`
	ArticleCount  int           `json:"articleCount"`
	IndexCount    int           `json:"indexCount"` //主页访问数量
	CreateAt      time.Time     `json:"createAt"`
	LastLoginDate *time.Time    `json:"lastLoginDate"`
	Role          enum.RoleType `json:"role"` // 角色
}

func (UserApi) UserListView(c *gin.Context) {
	cr := middleware.GetBind[UserListRequest](c)

	_list, _, _ := common.ListQuery(models.UserModel{}, common.Options{
		Likes:    []string{"nickname", "username"},
		Preloads: []string{"ArticleList", "LoginList"},
		PageInfo: cr.PageInfo,
	})

	var list = make([]UserListRespone, 0)
	for _, v := range _list {
		item := UserListRespone{
			ID:           v.ID,
			Nickname:     v.Nickname,
			Username:     v.Username,
			Avatar:       v.Avatar,
			IP:           v.Ip,
			Addr:         v.Addr,
			ArticleCount: len(v.ArticleList),
			IndexCount:   1000,
			CreateAt:     v.CreatedAt,
			Role:         v.Role,
		}
		if len(v.LoginList) > 0 {
			item.LastLoginDate = &v.LoginList[len(v.LoginList)-1].CreatedAt
		}
		list = append(list, item)
	}

	res.OkWithList(list, len(list), c)
}

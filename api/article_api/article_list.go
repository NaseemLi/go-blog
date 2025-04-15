package articleapi

import (
	"goblog/common"
	"goblog/common/res"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type ArticleListRequest struct {
	common.PageInfo
	Type       int8  `form:"type" binding:"required,oneof=1 2 3"` // 1 用户查别人的  2 查自己的  3 管理员查
	UserID     uint  `form:"userID"`
	CategoryID *uint `form:"categoryID"`
	Status     int   `form:"status"`
}

type ArticleListResponse struct {
	models.ArticleModel
	UserTop  bool `json:"userTop"`  //是否为用户置顶
	AdminTop bool `json:"adminTop"` //是否为管理员置顶
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	cr := middleware.GetBind[ArticleListRequest](c)

	switch cr.Type {
	case 1:
		//查别人用户 ID 必填
		if cr.UserID == 0 {
			res.FailWithMsg("用户ID必填", c)
			return
		}
		if cr.Page > 2 || cr.Limit > 10 {
			res.FailWithMsg("非登录用户,查询更多请登录", c)
			return
		}
		cr.Status = 0
	case 2:
		//查自己
		claims, err := jwts.ParseTokenByGin(c)
		if err != nil {
			res.FailWithMsg("请登录", c)
			return
		}
		cr.UserID = claims.UserID
	case 3:
		claims, err := jwts.ParseTokenByGin(c)
		if !(err == nil && claims.Role == enum.AdminRole) {
			res.FailWithMsg("角色错误", c)
			return
		}
	}
	_list, count, _ := common.ListQuery(models.ArticleModel{
		UserID:     cr.UserID,
		CategoryID: cr.CategoryID,
		Status:     enum.ArticleStatusType(cr.Status),
	}, common.Options{
		Likes:    []string{"title"},
		PageInfo: cr.PageInfo,
	})

	var list = make([]ArticleListResponse, 0)
	for _, model := range _list {
		model.Content = ""
		list = append(list, ArticleListResponse{
			ArticleModel: model,
		})
	}

	res.OkWithList(list, count, c)
}

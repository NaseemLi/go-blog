package articleapi

import (
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	"goblog/utils/jwts"

	"github.com/gin-gonic/gin"
)

type CategoryCreateRequest struct {
	ID    uint   `json:"id"`
	Title string `json:"title" binding:"required,max=32"`
}

func (ArticleApi) CategoryCreateView(c *gin.Context) {
	cr := middleware.GetBind[CategoryCreateRequest](c)

	claims := jwts.GetClaims(c)
	var model models.CategoryModel
	if cr.ID == 0 {
		//创建
		err := global.DB.Take(&model, "user_id = ? and title = ?", claims.UserID, cr.Title).Error
		if err == nil {
			res.FailWithMsg("分类名称重复", c)
			return
		}
		err = global.DB.Create(&models.CategoryModel{
			UserID: claims.UserID,
			Title:  cr.Title}).Error
		if err != nil {
			res.FailWithMsg("创建分类失败", c)
			return
		}
		res.OkWithMsg("创建分类成功", c)
		return
	}

	//更新
	err := global.DB.Take(&model, "user_id = ? and id = ?", claims.UserID, cr.ID).Error
	if err != nil {
		res.FailWithMsg("分类不存在", c)
		return
	}

	err = global.DB.Model(&model).Update("title", cr.Title).Error
	if err != nil {
		res.FailWithMsg("更新分类失败", c)
		return
	}
	res.OkWithMsg("更新分类成功", c)
}

type CategoryListRequest struct {
	common.PageInfo
	Type   int8 `form:"type" binding:"required,oneof=1 2 3"` //1.自己 2.别人 3.后台
	UserID uint `form:"userID"`                              // User ID for filtering
}

type CategoryCreateResponse struct {
	models.CategoryModel
	ArticleCount int    `json:"articleCount"`
	Nickname     string `json:"nickname,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
}

func (ArticleApi) CategoryListView(c *gin.Context) {
	cr := middleware.GetBind[CategoryListRequest](c)

	var preload = []string{"ArticleList"}
	switch cr.Type {
	case 1:
		claims, err := jwts.ParseTokenByGin(c)
		if err != nil {
			res.FailWithMsg("token解析失败", c)
			return
		}
		cr.UserID = claims.UserID
	case 2:

	case 3:
		claims, err := jwts.ParseTokenByGin(c)
		if err != nil {
			res.FailWithMsg("token解析失败", c)
			return
		}
		if claims.Role != enum.AdminRole {
			res.FailWithMsg("无权限访问", c)
			return
		}
		preload = append(preload, "UserModel")
	}
	_list, count, _ := common.ListQuery(models.CategoryModel{
		UserID: cr.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Likes:    []string{"title"},
		Preloads: preload,
	})

	var list = make([]CategoryCreateResponse, 0)
	for _, model := range _list {
		list = append(list, CategoryCreateResponse{
			CategoryModel: model,
			ArticleCount:  len(model.ArticleList),
			Nickname:      model.UserModel.Nickname,
			Avatar:        model.UserModel.Avatar,
		})
	}

	res.OkWithList(list, count, c)
}

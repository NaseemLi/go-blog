package articleapi

// ?逻辑
import (
	"fmt"
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	redisarticle "goblog/service/redis_service/redis_article"
	"goblog/utils/jwts"
	"goblog/utils/sql"
	"log"

	"github.com/gin-gonic/gin"
)

type ArticleListRequest struct {
	common.PageInfo
	Type       int8  `form:"type" binding:"required,oneof=1 2 3"` // 1 用户查别人的  2 查自己的  3 管理员查
	UserID     uint  `form:"userID"`
	CategoryID *uint `form:"categoryID"`
	Status     int   `form:"status"`
	CollectID  *uint `form:"collectID"`
}

type ArticleListResponse struct {
	models.ArticleModel
	UserTop       bool    `json:"userTop"`  //是否为用户置顶
	AdminTop      bool    `json:"adminTop"` //
	CategoryTitle *string `json:"categoryTitle"`
	UserNickname  string  `json:"userNickname"`
	UserAvatar    string  `json:"userAvatar"`
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	cr := middleware.GetBind[ArticleListRequest](c)

	var topArticleIDList []uint

	var orderColumnMap = map[string]bool{
		"look_count desc":    true,
		"digg_count desc":    true,
		"comment_count desc": true,
		"collect_count desc": true,
		"look_count asc":     true,
		"digg_count asc":     true,
		"comment_count asc":  true,
		"collect_count asc":  true,
	}

	switch cr.Type {
	case 1:
		if cr.UserID == 0 {
			res.FailWithMsg("用户ID必填", c)
			return
		}
		if cr.Page > 2 || cr.Limit > 10 {
			res.FailWithMsg("非登录用户,查询更多请登录", c)
			return
		}
		cr.Status = 0
		cr.Order = ""
		if cr.CollectID != nil && *cr.CollectID != 0 {
			// 如果传了收藏夹id,也要看人
			if cr.UserID == 0 {
				res.FailWithMsg("请传入用户id", c)
				return
			}

			var userConf models.UserConfModel
			err := global.DB.Take(&userConf, "id = ?", cr.UserID).Error
			if err != nil {
				res.FailWithMsg("用户不存在", c)
				return
			}

			if userConf.OpenCollect {
				res.FailWithMsg("用户未开放收藏夹", c)
				return
			}
		}
	case 2:
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

	if cr.Order != "" {
		_, ok := orderColumnMap[cr.Order]
		if !ok {
			res.FailWithMsg("不支持的排序方式", c)
			return
		}
	}

	var userTopMap = map[uint]bool{}
	var AdminTopMap = map[uint]bool{}

	if cr.UserID != 0 {
		var UserTopArticleList []models.UserTopArticleModel
		result := global.DB.Preload("UserModel").Order("created_at desc").Find(&UserTopArticleList, "user_id = ?", cr.UserID)

		if result.Error != nil {
			log.Printf("查询用户置顶文章失败: %v", result.Error)
			return
		}

		for _, i2 := range UserTopArticleList {
			topArticleIDList = append(topArticleIDList, i2.ArticleID)

			if i2.UserModel.ID != 0 && i2.UserModel.Role == enum.AdminRole {
				AdminTopMap[i2.ArticleID] = true
			}
			userTopMap[i2.ArticleID] = true
		}
	}

	var options = common.Options{
		Likes:        []string{"title"},
		PageInfo:     cr.PageInfo,
		DefaultOrder: "created_at desc",
		Preloads:     []string{"UserModel", "CategoryModel"},
	}

	if cr.Order != "" {
		// 用户指定排序字段优先
		options.DefaultOrder = fmt.Sprintf("%s, created_at desc", cr.Order)
	} else if len(topArticleIDList) > 0 {
		// 用户或管理员有置顶文章时按 FIELD 排序
		options.DefaultOrder = fmt.Sprintf("%s, created_at desc", sql.ConvertSliceOrderSql(topArticleIDList))
	} else {
		// 默认排序
		options.DefaultOrder = "created_at desc"
	}

	_list, count, _ := common.ListQuery(models.ArticleModel{
		UserID:     cr.UserID,
		CategoryID: cr.CategoryID,
		Status:     enum.ArticleStatusType(cr.Status),
	}, options)

	var list = make([]ArticleListResponse, 0)
	collectMap := redisarticle.GetAllCacheCollect()
	diggMap := redisarticle.GetAllCacheDigg()
	lookMap := redisarticle.GetAllCacheLook()
	commentMap := redisarticle.GetAllCacheComment()

	for _, model := range _list {
		model.Content = ""
		model.DiggCount = model.DiggCount + diggMap[model.ID]
		model.CollectCount = model.CollectCount + collectMap[model.ID]
		model.LookCount = model.LookCount + lookMap[model.ID]
		model.CommentCount = model.CommentCount + commentMap[model.ID]

		data := ArticleListResponse{
			ArticleModel: model,
			UserTop:      userTopMap[model.ID],
			AdminTop:     AdminTopMap[model.ID],
			UserNickname: model.UserModel.Nickname,
			UserAvatar:   model.UserModel.Avatar,
		}

		if model.CategoryModel != nil {
			data.CategoryTitle = &model.CategoryModel.Title
		}

		list = append(list, data)
	}

	res.OkWithList(list, count, c)
}

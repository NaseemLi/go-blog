package articleapi

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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ArticleLookRequest struct {
	ArticleID  uint `json:"articleID" binding:"required"`
	TimeSecond int  `json:"timeSecond"` // 读文章一共用了多久
}

func (ArticleApi) ArticleLookView(c *gin.Context) {
	cr := middleware.GetBind[ArticleLookRequest](c)

	claims, err := jwts.ParseTokenByGin(c)
	if err != nil {
		res.OkWithMsg("未登录", c)
		return
	}

	// 引入缓存
	// 当天这个用户请求这个文章之后，将用户id和文章id作为key存入缓存，在这里进行判断，如果存在就直接返回
	if redisarticle.GetUserArticleHistoryCache(cr.ArticleID, claims.UserID) {
		logrus.Info("在缓存内")
		res.OkWithMsg("成功", c)
		return
	}
	var article models.ArticleModel
	err = global.DB.Take(&article, "status = ? and id = ?", enum.ArticleStatusPublished, cr.ArticleID).Error
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	// 查这个文章今天有没有在足迹里面
	var history models.UserArticleLookHistoryModel
	err = global.DB.Take(&history,
		"user_id = ? and article_id = ? and created_at < ? and created_at > ?",
		claims.UserID, cr.ArticleID,
		time.Now().Format("2006-01-02 15:04:05"),
		time.Now().Format("2006-01-02")+" 00:00:00",
	).Error
	if err == nil {
		res.OkWithMsg("成功", c)
		return
	}

	err = global.DB.Create(&models.UserArticleLookHistoryModel{
		UserID:    claims.UserID,
		ArticleID: cr.ArticleID,
	}).Error
	if err != nil {
		res.FailWithMsg("失败", c)
		return
	}
	redisarticle.SetCacheLook(cr.ArticleID, true)
	redisarticle.SetUserArticleHistoryCache(cr.ArticleID, claims.UserID)
	res.OkWithMsg("成功", c)
}

type ArticleLookListRequest struct {
	common.PageInfo
	UserID uint `form:"userID"`
	Type   int8 `form:"type" binding:"required,oneof=1 2"`
}

type ArticleLookListResponse struct {
	ID        uint      `json:"id"`
	LookDate  time.Time `json:"lookDate"`
	Title     string    `json:"title"`
	Cover     string    `json:"cover"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	UserID    uint      `json:"userID"`
	ArticleID uint      `json:"articleID"`
}

func (ArticleApi) ArticleLookListView(c *gin.Context) {
	cr := middleware.GetBind[ArticleLookListRequest](c)
	claims := jwts.GetClaims(c)

	switch cr.Type {
	case 1:
		cr.UserID = claims.UserID
	}

	_list, count, _ := common.ListQuery(models.UserArticleLookHistoryModel{
		UserID: cr.UserID,
	}, common.Options{
		PageInfo: cr.PageInfo,
		Preloads: []string{"UserModel", "ArticleModel"},
	})

	var list = make([]ArticleLookListResponse, 0)
	for _, model := range _list {
		list = append(list, ArticleLookListResponse{
			ID:        model.ID,
			LookDate:  model.CreatedAt,
			Title:     model.ArticleModel.Title,
			Cover:     model.ArticleModel.Cover,
			Nickname:  model.UserModel.Nickname,
			Avatar:    model.UserModel.Avatar,
			UserID:    model.UserModel.ID,
			ArticleID: model.ArticleID,
		})
	}

	res.OkWithList(list, count, c)
}

func (ArticleApi) ArticleLookRemoveView(c *gin.Context) {
	cr := middleware.GetBind[models.RemoveRequest](c)

	claims := jwts.GetClaims(c)

	var list []models.UserArticleLookHistoryModel
	global.DB.Find(&list, "user_id = ? and id in ?", claims.UserID, cr.IDList)
	if len(list) == 0 {
		res.FailWithMsg("没有找到记录", c)
		return
	}

	global.DB.Delete(&list)

	res.OkWithMsg(fmt.Sprintf("删除足迹成功,共删除 %d 条", len(list)), c)
}

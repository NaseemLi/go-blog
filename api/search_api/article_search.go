package searchapi

import (
	"context"
	"encoding/json"
	"fmt"
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"goblog/models/enum"
	"goblog/utils/jwts"
	"goblog/utils/sql"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type ArticleSearchRequest struct {
	common.PageInfo
	Tag  string `form:"tag"`
	Type int8   `form:"type"` // 0 猜你喜欢  1 最新发布  2最多回复 3最多点赞 4最多收藏
}

type ArticleBaseInfo struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
}

type ArticleListResponse struct {
	models.ArticleModel
	AdminTop      bool    `json:"adminTop"` // 是否是管理员置顶
	CategoryTitle *string `json:"categoryTitle"`
	UserNickname  string  `json:"userNickname"`
	UserAvatar    string  `json:"userAvatar"`
}

func (SearchApi) ArticleSearchView(c *gin.Context) {
	var cr = middleware.GetBind[ArticleSearchRequest](c)

	var sortMap = map[int8]string{
		0: "_score",
		1: "created_at",
		2: "comment_count",
		3: "digg_count",
		4: "collect_count",
	}
	sortKey := sortMap[cr.Type]
	if sortKey == "" {
		res.FailWithMsg("搜索类型错误", c)
		return
	}

	topArticleIDList := getAdminTopArticleIDList()

	if global.ESClient == nil {
		//服务讲解,用户可能没有配置 es
		where := global.DB.Where("")
		if cr.Tag != "" {
			where = where.Where("tag_list like ?", fmt.Sprintf("%%%s%%", cr.Tag))
		}
		var articleTopMap = map[uint]bool{}
		for _, v := range topArticleIDList {
			articleTopMap[v] = true
		}

		sortMap = map[int8]string{
			0: "",
			1: "created_at desc",
			2: "comment_count desc",
			3: "digg_count desc",
			4: "collect_count desc",
		}
		sort, _ := sortMap[cr.Type]
		cr.Order = sort

		_list, count, _ := common.ListQuery(models.ArticleModel{}, common.Options{
			Preloads:     []string{"CategoryModel", "UserModel"},
			PageInfo:     cr.PageInfo,
			Likes:        []string{"title", "abstract", "tag_list"},
			DefaultOrder: sql.ConvertSliceOrderSql(topArticleIDList),
			Where:        where,
		})

		var list = make([]ArticleListResponse, 0)
		for _, model := range _list {
			item := ArticleListResponse{
				ArticleModel: model,
				AdminTop:     articleTopMap[model.ID],
				UserNickname: model.UserModel.Nickname,
				UserAvatar:   model.UserModel.Avatar,
			}
			if model.CategoryModel != nil {
				item.CategoryTitle = &model.CategoryModel.Title
			}
			list = append(list, item)
		}

		res.OkWithList(list, count, c)
		return
	}

	query := elastic.NewBoolQuery()
	if cr.Key != "" {
		query.Should(
			elastic.NewMatchQuery("title", cr.Key),
			elastic.NewMatchQuery("abstract", cr.Key),
			elastic.NewMatchQuery("content", cr.Key),
		)
	}
	if cr.Tag != "" {
		query.Must(
			elastic.NewTermQuery("tag_list", cr.Tag),
		)
	}

	// 只能查发布的文章
	query.Must(elastic.NewTermQuery("status", 3))

	var articleIDList []uint

	// 把管理员置顶的文章查出来
	var articleTopMap = map[uint]bool{}
	if len(topArticleIDList) > 0 {
		var topArticleIDListAny []interface{}
		for _, u := range topArticleIDList {
			topArticleIDListAny = append(topArticleIDListAny, u)
			articleTopMap[u] = true
			articleIDList = append(articleIDList, u)
		}
		query.Should(elastic.NewTermsQuery("id", topArticleIDListAny...).Boost(10))
	}

	if cr.Type == 0 {
		// 只有猜你喜欢，才会把用户喜欢的标签带入查询
		claims, err := jwts.ParseTokenByGin(c)
		if err == nil && claims != nil {
			// 用户登录了
			// 查用户感兴趣的分类
			var userConf models.UserConfModel
			err = global.DB.Take(&userConf, "user_id = ?", claims.UserID).Error
			if err != nil {
				res.FailWithMsg("用户配置不存在", c)
				return
			}
			if len(userConf.LikeTags) > 0 {
				tagQuery := elastic.NewBoolQuery()
				var tagAnyList []interface{}
				for _, tag := range userConf.LikeTags {
					tagAnyList = append(tagAnyList, tag)
				}
				tagQuery.Should(elastic.NewTermsQuery("tag_list", tagAnyList...))
				query.Must(tagQuery)
			}
		}
	}

	highlight := elastic.NewHighlight()
	highlight.Field("title")
	highlight.Field("abstract")

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Highlight(highlight).
		From(cr.GetOffset()).
		Size(cr.GetLimit()).
		Sort(sortKey, false).
		Do(context.Background())
	if err != nil {
		source, _ := query.Source()
		byteData, _ := json.Marshal(source)
		logrus.Errorf("查询失败 %s \n %s", err, string(byteData))
		res.FailWithMsg("查询失败", c)
		return
	}

	count := result.Hits.TotalHits.Value
	var searchArticleMap = map[uint]ArticleBaseInfo{}

	for _, hit := range result.Hits.Hits {

		var art ArticleBaseInfo
		err = json.Unmarshal(hit.Source, &art)
		if err != nil {
			logrus.Warnf("解析失败 %s  %s", err, string(hit.Source))
			continue
		}
		if hit.Score != nil {
			fmt.Println(*hit.Score, art.Title, art.ID)
		}
		if len(hit.Highlight["title"]) > 0 {
			art.Title = hit.Highlight["title"][0]
		}
		if len(hit.Highlight["abstract"]) > 0 {
			art.Abstract = hit.Highlight["abstract"][0]
		}

		searchArticleMap[art.ID] = art
		articleIDList = append(articleIDList, art.ID)
	}

	where := global.DB.Where("id in ?", articleIDList)

	_list, _, _ := common.ListQuery(models.ArticleModel{}, common.Options{
		Where:        where,
		Preloads:     []string{"CategoryModel", "UserModel"},
		DefaultOrder: sql.ConvertSliceOrderSql(articleIDList),
	})

	var list = make([]ArticleListResponse, 0)
	for _, model := range _list {
		item := ArticleListResponse{
			ArticleModel: model,
			AdminTop:     articleTopMap[model.ID],
			UserNickname: model.UserModel.Nickname,
			UserAvatar:   model.UserModel.Avatar,
		}
		if model.CategoryModel != nil {
			item.CategoryTitle = &model.CategoryModel.Title
		}
		item.Title = searchArticleMap[model.ID].Title
		item.Abstract = searchArticleMap[model.ID].Abstract
		list = append(list, item)
	}

	res.OkWithList(list, int(count), c)
}

func getAdminTopArticleIDList() (topArticleIDList []uint) {
	var userIDList []uint
	global.DB.Model(models.UserModel{}).Where("role = ?", enum.AdminRole).Select("id").Scan(&userIDList)
	global.DB.Model(models.UserTopArticleModel{}).Where("user_id in ?", userIDList).Select("article_id").Scan(&topArticleIDList)
	return
}

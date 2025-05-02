package searchapi

import (
	"context"
	"encoding/json"
	"goblog/common"
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	"goblog/models"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type AggType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
	} `json:"buckets"`
}

type Agg1Type struct {
	Value int `json:"value"`
}

type TagAggResponse struct {
	Tags         string `json:"tags"`
	ArticleCount int    `json:"articleCount"`
}

func (SearchApi) TagAggView(c *gin.Context) {
	var cr = middleware.GetBind[common.PageInfo](c)
	var list = make([]TagAggResponse, 0)

	if global.ESClient == nil {
		var articleList []models.ArticleModel
		global.DB.Find(&articleList, "tag_list != ''")
		var tagMap = make(map[string]int)
		for _, v := range articleList {
			for _, tag := range v.TagList {
				count, ok := tagMap[tag]
				if !ok {
					tagMap[tag] = 1
					continue
				}
				tagMap[tag] = count + 1
			}
		}
		for tag, count := range tagMap {
			list = append(list, TagAggResponse{
				Tags:         tag,
				ArticleCount: count,
			})
		}
		sort.Slice(list, func(i, j int) bool {
			return list[i].ArticleCount > list[j].ArticleCount
		})

		res.OkWithList(list, len(list), c)
		return
	}

	agg := elastic.NewTermsAggregation().Field("tag_list")
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(cr.GetOffset()).Size(cr.Limit))
	// agg.SubAggregation("sum", elastic.NewSumAggregation().Field("tag_list"))

	query := elastic.NewBoolQuery()
	query.MustNot(elastic.NewTermQuery("tag_list", ""))
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Aggregation("tags1", elastic.NewCardinalityAggregation().Field("tag_list")).
		Size(0).
		Do(context.Background())

	if err != nil {
		logrus.Errorf("es查询失败: %v", err)
		res.FailWithMsg("查询失败，请稍后再试", c)
		return
	}
	var t AggType
	var val = result.Aggregations["tags"]
	err = json.Unmarshal(val, &t)
	if err != nil {
		logrus.Errorf("es反序列化失败: %v", err)
		res.FailWithMsg("查询失败，请稍后再试", c)
		return
	}

	var co Agg1Type
	json.Unmarshal(result.Aggregations["tags1"], &co)

	for _, v := range t.Buckets {
		list = append(list, TagAggResponse{
			Tags:         v.Key,
			ArticleCount: v.DocCount,
		})
	}

	res.OkWithList(list, co.Value, c)
}

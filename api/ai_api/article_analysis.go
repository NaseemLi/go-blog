package aiapi

import (
	"goblog/common/res"
	"goblog/global"
	"goblog/middleware"
	aiservice "goblog/service/ai_service"

	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ArticleAnalysisRequest struct {
	Content string `json:"content" binding:"required"`
}

type ArticleAnalysisResponse struct {
	Title    string   `json:"title"`
	Abstract string   `json:"abstract"`
	Category string   `json:"category"`
	Tag      []string `json:"tag"`
}

func (AiApi) ArticleAnalysisView(c *gin.Context) {
	cr := middleware.GetBind[ArticleAnalysisRequest](c)

	if !global.Config.Ai.Enable {
		res.FailWithMsg("AI功能未开启", c)
		return
	}

	msg, err := aiservice.Chat(cr.Content)
	if err != nil {
		logrus.Errorf("Ai分析错误%v %s", err, cr.Content)
		res.FailWithMsg("Ai分析错误", c)
		return
	}
	var response ArticleAnalysisResponse
	err = json.Unmarshal([]byte(msg), &response)
	if err != nil {
		logrus.Errorf("Ai解析失败%v %s", err, msg)
		res.FailWithMsg("Ai解析失败", c)
		return
	}

	res.OkWithData(response, c)
}

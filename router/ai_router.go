package router

import (
	"goblog/api"
	aiapi "goblog/api/ai_api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func AiRouter(r *gin.RouterGroup) {
	app := api.App.AiApi
	r.POST("ai/analysis", middleware.AuthMiddelware, middleware.BindJsonMiddleware[aiapi.ArticleAnalysisRequest], app.ArticleAnalysisView)
}

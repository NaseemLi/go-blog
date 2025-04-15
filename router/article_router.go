package router

import (
	"goblog/api"
	articleapi "goblog/api/article_api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func ArticleRouter(r *gin.RouterGroup) {
	app := api.App.ArticleApi
	r.POST("article", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.ArticleCreateRequest], app.ArticleCreateView)
}

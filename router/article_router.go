package router

import (
	"goblog/api"
	articleapi "goblog/api/article_api"
	"goblog/middleware"
	"goblog/models"

	"github.com/gin-gonic/gin"
)

func ArticleRouter(r *gin.RouterGroup) {
	app := api.App.ArticleApi
	r.POST("article", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.ArticleCreateRequest], app.ArticleCreateView)
	r.PUT("article", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.ArticleUpdateRequest], app.ArticleUpdateView)
	r.GET("article", middleware.BindQueryMiddleware[articleapi.ArticleListRequest], app.ArticleListView)
	r.GET("article/:id", middleware.BindUriMiddleware[models.IDRequest], app.ArticleDetailView)
}

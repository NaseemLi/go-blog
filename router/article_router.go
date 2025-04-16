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
	r.POST("article/examine", middleware.AdminMiddelware, middleware.BindJsonMiddleware[articleapi.ArticleExamineRequest], app.ArticleExamineView)
	r.GET("article/digg/:id", middleware.AuthMiddelware, middleware.BindUriMiddleware[models.IDRequest], app.ArticleDiggView)
	r.POST("article/collect", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.ArticleCollectRequest], app.ArticleCollectView)
	r.POST("article/look", middleware.BindJsonMiddleware[articleapi.ArticleLookRequest], app.ArticleLookView)
}

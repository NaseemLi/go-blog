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
	r.DELETE("article/:id", middleware.AuthMiddelware, middleware.BindUriMiddleware[models.IDRequest], app.ArticleRemoveUserView)
	r.DELETE("article", middleware.AdminMiddelware, middleware.BindJsonMiddleware[models.RemoveRequest], app.ArticleRemoveAdminView)

	r.GET("article/history", middleware.AuthMiddelware, middleware.BindQueryMiddleware[articleapi.ArticleLookListRequest], app.ArticleLookListView)
	r.DELETE("article/history", middleware.AuthMiddelware, middleware.BindJsonMiddleware[models.RemoveRequest], app.ArticleLookRemoveView)

	r.POST("category", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.CategoryCreateRequest], app.CategoryCreateView)
	r.GET("category", middleware.BindQueryMiddleware[articleapi.CategoryListRequest], app.CategoryListView)
	r.DELETE("category", middleware.AuthMiddelware, middleware.BindJsonMiddleware[models.RemoveRequest], app.CategoryRemoveView)
	// 收藏夹的接口（收藏夹 CRUD）
	r.POST("collect", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.CollectCreateRequest], app.CollectCreateView)
	r.GET("collect", middleware.BindQueryMiddleware[articleapi.CollectListRequest], app.CollectListView)
	r.DELETE("collect", middleware.AuthMiddelware, middleware.BindJsonMiddleware[models.RemoveRequest], app.CollectRemoveView)
	// 文章的收藏夹接口
	r.GET("category/options", middleware.AuthMiddelware, app.CategoryOptionsView)
}

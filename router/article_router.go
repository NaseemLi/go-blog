package router

import (
	"goblog/api"
	articleapi "goblog/api/article_api"
	"goblog/common"
	"goblog/middleware"
	"goblog/models"

	"github.com/gin-gonic/gin"
)

func ArticleRouter(r *gin.RouterGroup) {
	app := api.App.ArticleApi

	// ------------------ 文章接口 ------------------
	r.POST("article", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.ArticleCreateRequest], app.ArticleCreateView)
	r.PUT("article", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.ArticleUpdateRequest], app.ArticleUpdateView)
	r.GET("article", middleware.BindQueryMiddleware[articleapi.ArticleListRequest], app.ArticleListView)
	r.GET("article/:id", middleware.BindUriMiddleware[models.IDRequest], app.ArticleDetailView)
	r.DELETE("article/:id", middleware.AuthMiddelware, middleware.BindUriMiddleware[models.IDRequest], app.ArticleRemoveUserView)
	r.DELETE("article", middleware.AdminMiddelware, middleware.BindJsonMiddleware[models.RemoveRequest], app.ArticleRemoveAdminView)

	// ------------------ 审核相关 ------------------
	r.POST("article/examine", middleware.AdminMiddelware, middleware.BindJsonMiddleware[articleapi.ArticleExamineRequest], app.ArticleExamineView)

	// ------------------ 点赞 / 收藏 / 阅读 ------------------
	r.GET("article/digg/:id", middleware.AuthMiddelware, middleware.BindUriMiddleware[models.IDRequest], app.ArticleDiggView)
	r.POST("article/collect", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.ArticleCollectRequest], app.ArticleCollectView)
	r.DELETE("article/collect", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.ArticleCollectPatchRemoveRequest], app.ArticleCollectPatchRemoveView)
	r.POST("article/look", middleware.BindJsonMiddleware[articleapi.ArticleLookRequest], app.ArticleLookView)

	// ------------------ 阅读历史 ------------------
	r.GET("article/history", middleware.AuthMiddelware, middleware.BindQueryMiddleware[articleapi.ArticleLookListRequest], app.ArticleLookListView)
	r.DELETE("article/history", middleware.AuthMiddelware, middleware.BindJsonMiddleware[models.RemoveRequest], app.ArticleLookRemoveView)

	// ------------------ 分类管理 ------------------
	r.POST("category", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.CategoryCreateRequest], app.CategoryCreateView)
	r.GET("category", middleware.BindQueryMiddleware[articleapi.CategoryListRequest], app.CategoryListView)
	r.DELETE("category", middleware.AuthMiddelware, middleware.BindJsonMiddleware[models.RemoveRequest], app.CategoryRemoveView)
	r.GET("category/options", middleware.AuthMiddelware, app.CategoryOptionsView)

	// ------------------ 收藏夹（收藏夹 CRUD） ------------------
	r.POST("collect", middleware.AuthMiddelware, middleware.BindJsonMiddleware[articleapi.CollectCreateRequest], app.CollectCreateView)
	r.GET("collect", middleware.BindQueryMiddleware[articleapi.CollectListRequest], app.CollectListView)
	r.DELETE("collect", middleware.AuthMiddelware, middleware.BindJsonMiddleware[models.RemoveRequest], app.CollectRemoveView)

	// ------------------ 智能推荐 ------------------
	r.GET("article/auth_recommend", middleware.BindQueryMiddleware[common.PageInfo], app.AuthRecommentView)
	r.GET("article/article_recommend", middleware.BindQueryMiddleware[common.PageInfo], app.ArticleRecommentView)

	// ------------------ 标签配置 ------------------
	r.GET("article/tag/options", middleware.AuthMiddelware, app.ArticleTagOptions)
}

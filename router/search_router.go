package router

import (
	"goblog/api"
	searchapi "goblog/api/search_api"
	"goblog/common"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func SearchRouter(r *gin.RouterGroup) {
	app := api.App.SearchApi
	r.GET("search/article", middleware.BindQueryMiddleware[searchapi.ArticleSearchRequest], app.ArticleSearchView)
	r.GET("search/tags", middleware.BindQueryMiddleware[common.PageInfo], app.TagAggView)
	r.GET("search/text", middleware.BindQueryMiddleware[searchapi.TextSearchRequest], app.TextSearchView)
}

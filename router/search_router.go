package router

import (
	"goblog/api"
	searchapi "goblog/api/search_api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func SearchRouter(r *gin.RouterGroup) {
	app := api.App.SearchApi
	r.GET("search/article", middleware.BindQueryMiddleware[searchapi.ArticleSearchRequest], app.ArticleSearchView)
}

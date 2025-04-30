package router

import (
	"goblog/api"
	dataapi "goblog/api/data_api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func DataRouter(r *gin.RouterGroup) {
	app := api.App.DataApi
	r.GET("data/sum", middleware.AdminMiddelware, app.SumView)
	r.GET("data/article", middleware.AdminMiddelware, app.ArticleDataView)
	r.GET("data/computer", middleware.AdminMiddelware, app.ComputerDataView)
	r.GET("data/growth", middleware.AdminMiddelware, middleware.BindQueryMiddleware[dataapi.GrowthDataRequest], app.GrowthDataView)
}

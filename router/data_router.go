package router

import (
	"goblog/api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func DataRouter(r *gin.RouterGroup) {
	app := api.App.DataApi
	r.GET("data/sum", middleware.AdminMiddelware, app.SumView)
}

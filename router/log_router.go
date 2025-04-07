package router

import (
	"goblog/api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func LogRouter(r *gin.RouterGroup) {
	app := api.App.LogApi
	r.Use(middleware.AdminMiddelware)
	r.GET("logs", app.LogListView)
	r.GET("logs/:id", app.LogReadView)
	r.DELETE("logs", app.LogRemoveView)
}

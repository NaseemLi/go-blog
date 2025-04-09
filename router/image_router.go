package router

import (
	"goblog/api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	app := api.App.ImageApi
	r.POST("images", middleware.AuthMiddelware, app.ImageUploadView)
	r.GET("images", middleware.AdminMiddelware, app.ImageListView)
	r.DELETE("images", middleware.AdminMiddelware, app.ImageRemoveView)
}

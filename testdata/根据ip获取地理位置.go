package router

import (
	"goblog/api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.Siteapi
	r.POST("images", middleware.AuthMiddelware, app.ImageUploadView)
}

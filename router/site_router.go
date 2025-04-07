package router

import (
	"goblog/api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.Siteapi
	r.GET("site", app.SiteInfoView)
	r.PUT("site", middleware.AdminMiddelware, app.SiteUpdateView)
}

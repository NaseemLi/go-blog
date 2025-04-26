package router

import (
	"goblog/api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.Siteapi
	r.GET("site/qq_url", app.SiteInfoQQView)
	r.GET("site/ai_info", app.SiteInfoAiView)
	r.GET("site/:name", app.SiteInfoView)

	r.PUT("site/:name", middleware.AdminMiddelware, app.SiteUpdateView)
}

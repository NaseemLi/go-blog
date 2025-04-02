package router

import (
	"goblog/api"

	"github.com/gin-gonic/gin"
)

func SiteRouter(r *gin.RouterGroup) {
	app := api.App.Siteapi
	r.GET("site", app.SiteInfoView)
	r.PUT("site", app.SiteUpdateView)
}

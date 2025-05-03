package router

import (
	"goblog/api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func BannerRouter(r *gin.RouterGroup) {
	app := api.App.BannerApi
	r.GET("banner", middleware.CacheMiddleware(middleware.NewBannerCacheOptions()), app.BannerListView)
	r.POST("banner", middleware.AdminMiddelware, app.BannerCreateView)
	r.PUT("banner/:id", middleware.AdminMiddelware, app.BannerUpdateView)
	r.DELETE("banner", middleware.AdminMiddelware, app.BannerRemoveView)

}

package router

import (
	"goblog/api"
	sitemsgapi "goblog/api/site_msg_api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func SiteMsgRouter(r *gin.RouterGroup) {
	app := api.App.SiteMsgApi
	r.GET("site_msg", middleware.AuthMiddelware, middleware.BindQueryMiddleware[sitemsgapi.SiteMsgListRequest], app.SiteMsgListView)
	r.GET("site_msg/conf", middleware.AuthMiddelware, app.UserSiteMessageConfView)
	r.PUT("site_msg/conf", middleware.AuthMiddelware, middleware.BindJsonMiddleware[sitemsgapi.UserMessageConfUpdateRequest], app.UserSiteMessageConfUpdateView)
	r.POST("site_msg", middleware.AuthMiddelware, middleware.BindJsonMiddleware[sitemsgapi.SiteMsgReadRequest], app.SiteMsgReadView)
	r.DELETE("site_msg", middleware.AuthMiddelware, middleware.BindJsonMiddleware[sitemsgapi.SiteMsgRemoveRequest], app.SiteMsgRemoveView)
	r.GET("site_msg/user", middleware.AuthMiddelware, app.UserMsgView)
}

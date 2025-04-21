package router

import (
	"goblog/api"
	globalnotificationapi "goblog/api/global_notification_api"
	"goblog/middleware"
	"goblog/models"

	"github.com/gin-gonic/gin"
)

func GlobalNotificationRouter(r *gin.RouterGroup) {
	app := api.App.GlobalNotificationApi
	r.POST("global_notification", middleware.AuthMiddelware, middleware.BindJsonMiddleware[globalnotificationapi.CreateRequest], app.CreateView)
	r.GET("global_notification", middleware.AuthMiddelware, middleware.BindQueryMiddleware[globalnotificationapi.ListRequest], app.ListView)
	r.DELETE("global_notification", middleware.AdminMiddelware, middleware.BindJsonMiddleware[models.RemoveRequest], app.RemoveView)
	r.POST("global_notification/user", middleware.AuthMiddelware, middleware.BindJsonMiddleware[globalnotificationapi.UserMsgActionViewRequest], app.UserMsgActionView)
}

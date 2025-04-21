package router

import (
	"goblog/api"
	focusapi "goblog/api/focus_api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func FocusRouter(r *gin.RouterGroup) {
	app := api.App.FocusApi
	r.POST("focus", middleware.AuthMiddelware, middleware.BindJsonMiddleware[focusapi.FocusUserRequest], app.FocusUserView)
	r.GET("focus", middleware.AuthMiddelware, middleware.BindQueryMiddleware[focusapi.FocusUserListRequest], app.FocusUserListView)
}

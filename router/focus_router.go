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
	r.DELETE("focus", middleware.AuthMiddelware, middleware.BindJsonMiddleware[focusapi.UnFocusUserRequest], app.UnFocusUserView)
	r.GET("focus/my_focus", middleware.BindQueryMiddleware[focusapi.FocusUserListRequest], app.FocusUserListView)
	r.GET("focus/my_fans", middleware.BindQueryMiddleware[focusapi.FocusUserListRequest], app.FansUserListView)
}

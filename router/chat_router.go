package router

import (
	"goblog/api"
	chatapi "goblog/api/chat_api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func ChatRouter(r *gin.RouterGroup) {
	app := api.App.ChatApi
	r.GET("chat", middleware.AuthMiddelware, middleware.BindQueryMiddleware[chatapi.ChatRecordRequest], app.ChatRecordView)
	r.GET("chat/session", middleware.AuthMiddelware, middleware.BindQueryMiddleware[chatapi.SessionListRequest], app.SessionListView)
}

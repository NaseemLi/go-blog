package router

import (
	"goblog/api"
	chatapi "goblog/api/chat_api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func ChatRouter(r *gin.RouterGroup) {
	app := api.App.ChatApi
	r.GET("chat/record", middleware.AuthMiddelware, middleware.BindQueryMiddleware[chatapi.ChatRecordRequest], app.ChatRecordView)
}

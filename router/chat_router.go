package router

import (
	"goblog/api"
	chatapi "goblog/api/chat_api"
	"goblog/middleware"
	"goblog/models"

	"github.com/gin-gonic/gin"
)

func ChatRouter(r *gin.RouterGroup) {
	app := api.App.ChatApi
	r.GET("chat", middleware.AuthMiddelware, middleware.BindQueryMiddleware[chatapi.ChatRecordRequest], app.ChatRecordView)
	r.GET("chat/session", middleware.AuthMiddelware, middleware.BindQueryMiddleware[chatapi.SessionListRequest], app.SessionListView)
	r.DELETE("chat", middleware.AuthMiddelware, middleware.BindJsonMiddleware[models.RemoveRequest], app.UserChatDeleteView)
	r.DELETE("chat/user/:id", middleware.AuthMiddelware, middleware.BindUriMiddleware[models.IDRequest], app.UserChatDeleteByUserView)
	r.POST("chat/read/:id", middleware.AuthMiddelware, middleware.BindUriMiddleware[models.IDRequest], app.UserChatReadView)
	r.GET("chat/ws", app.ChatView) //不要加任何验证
}

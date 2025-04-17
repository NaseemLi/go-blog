package router

import (
	"goblog/api"
	commentapi "goblog/api/comment_api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.RouterGroup) {
	app := api.App.CommentApi
	r.POST("comment", middleware.AuthMiddelware, middleware.BindJsonMiddleware[commentapi.CommentCreateRequest], app.CommentCreateView)
}

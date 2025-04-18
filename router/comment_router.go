package router

import (
	"goblog/api"
	commentapi "goblog/api/comment_api"
	"goblog/middleware"
	"goblog/models"

	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.RouterGroup) {
	app := api.App.CommentApi
	r.POST("comment", middleware.AuthMiddelware, middleware.BindJsonMiddleware[commentapi.CommentCreateRequest], app.CommentCreateView)
	r.GET("comment/tree/:id", middleware.BindUriMiddleware[models.IDRequest], app.CommentTreeView)
	r.GET("comment", middleware.AuthMiddelware, middleware.BindQueryMiddleware[commentapi.CommentListRequest], app.CommentListView)
	r.DELETE("comment/:id", middleware.AuthMiddelware, middleware.BindUriMiddleware[models.IDRequest], app.CommentRemoveView)
}

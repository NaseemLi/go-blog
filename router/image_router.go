package router

import (
	"goblog/api"
	imageapi "goblog/api/image_api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func ImageRouter(r *gin.RouterGroup) {
	app := api.App.ImageApi
	r.POST("images", middleware.AuthMiddelware, app.ImageUploadView)
	r.POST("images/qiniu", middleware.AuthMiddelware, app.QiNiuGenToken)
	r.GET("images", middleware.AdminMiddelware, app.ImageListView)
	r.DELETE("images", middleware.AdminMiddelware, app.ImageRemoveView)
	r.POST("images/transfer_deposit", middleware.AdminMiddelware, middleware.BindJsonMiddleware[imageapi.TransferDepositRequest], app.TransferDepositView)
}

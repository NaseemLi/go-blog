package router

import (
	"goblog/api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	app := api.App.UserApi
	r.POST("user/send_email", middleware.CaptchaMiddleware, app.SendEmailView)
	r.POST("user/email", app.RegisterEmailView)
	r.POST("user/qq", app.QQLoginView)
	r.POST("user/login", middleware.CaptchaMiddleware, app.PwdLoginApi)
	r.GET("user/detail", middleware.AuthMiddelware, app.UserDetailView)
	r.GET("user/base", app.UserBaseInfoView)
	r.GET("user/login", middleware.AuthMiddelware, app.UserLoginListView)
	r.PUT("user/password", middleware.AuthMiddelware, app.UpdatePasswordView)
}

package router

import (
	"goblog/api"
	userapi "goblog/api/user_api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	app := api.App.UserApi
	r.POST("user/send_email", middleware.CaptchaMiddleware, app.SendEmailView)
	r.POST("user/email", app.RegisterEmailView)
	r.POST("user/qq", app.QQLoginView)
	r.POST("user/login", middleware.CaptchaMiddleware, middleware.BindJsonMiddleware[userapi.PwdLoginrequest], app.PwdLoginApi)
	r.GET("user/detail", middleware.AuthMiddelware, app.UserDetailView)
	r.GET("user", middleware.AdminMiddelware, middleware.BindQueryMiddleware[userapi.UserListRequest], app.UserListView)
	r.GET("user/base", app.UserBaseInfoView)
	r.GET("user/login", middleware.AuthMiddelware, app.UserLoginListView)
	r.PUT("user/password", middleware.AuthMiddelware, app.UpdatePasswordView)
	r.PUT("user/password/reset", middleware.CaptchaMiddleware, app.ResetPasswordView)
	r.PUT("user/email/bind", middleware.CaptchaMiddleware, middleware.AuthMiddelware, app.BindEmailView)
	r.PUT("user", middleware.AuthMiddelware, app.UserInfoUpdateView)
	r.PUT("user/admin", middleware.AdminMiddelware, app.AdminInfoUpdateView)
	r.DELETE("user/logout", middleware.AuthMiddelware, app.LogOutView)
	r.POST("user/article/top", middleware.AuthMiddelware, middleware.BindJsonMiddleware[userapi.UserArticleTopRequest], app.UserArticleTopView)
}

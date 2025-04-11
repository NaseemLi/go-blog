package router

import (
	"goblog/api"
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func CaptchaRouter(r *gin.RouterGroup) {
	app := api.App.Captcha
	r.GET("captcha", middleware.AuthMiddelware, app.CaptchaView)
}

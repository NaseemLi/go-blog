package api

import (
	articleapi "goblog/api/article_api"
	bannerapi "goblog/api/banner_api"
	captchaapi "goblog/api/captcha_api"
	commentapi "goblog/api/comment_api"
	imageapi "goblog/api/image_api"
	logapi "goblog/api/log_api"
	siteapi "goblog/api/site_api"
	userapi "goblog/api/user_api"
)

type Api struct {
	Siteapi    siteapi.Siteapi
	LogApi     logapi.LogApi
	ImageApi   imageapi.ImageApi
	BannerApi  bannerapi.BannerApi
	Captcha    captchaapi.CaptchaApi
	UserApi    userapi.UserApi
	ArticleApi articleapi.ArticleApi
	CommentApi commentapi.CommentApi
}

var App = Api{}

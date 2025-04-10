package api

import (
	bannerapi "goblog/api/banner_api"
	imageapi "goblog/api/image_api"
	logapi "goblog/api/log_api"
	siteapi "goblog/api/site_api"
)

type Api struct {
	Siteapi   siteapi.Siteapi
	LogApi    logapi.LogApi
	ImageApi  imageapi.ImageApi
	BannerApi bannerapi.BannerApi
}

var App = Api{}

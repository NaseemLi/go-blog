package api

import (
	imageapi "goblog/api/image_api"
	logapi "goblog/api/log_api"
	siteapi "goblog/api/site_api"
)

type Api struct {
	Siteapi  siteapi.Siteapi
	LogApi   logapi.LogApi
	ImageApi imageapi.ImageApi
}

var App = Api{}

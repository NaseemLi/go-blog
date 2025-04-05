package api

import (
	logapi "goblog/api/log_api"
	siteapi "goblog/api/site_api"
)

type Api struct {
	Siteapi siteapi.Siteapi
	LogApi  logapi.LogApi
}

var App = Api{}

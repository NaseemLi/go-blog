package api

import siteapi "goblog/api/site_api"

type Api struct {
	Siteapi siteapi.Siteapi
}

var App = Api{}

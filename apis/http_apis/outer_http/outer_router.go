package outer_http

import (
	"myweb/apis/http_apis"
	router_tool "myweb/tools/router"
)

var OuterUrlList = []router_tool.RouterStruct{
	// user apis
	{Method: http_apis.Method.POST, Path: "/user/login", Handler: ApiUserLogin},
	{Method: http_apis.Method.GET, Path: "/user/check_token", Handler: ApiCheckToken},
	{Method: http_apis.Method.GET, Path: "/user/info", Handler: ApiGetUserInfo},

	// villain apis
	{Method: http_apis.Method.GET, Path: "/villain/info", Handler: ApiVillainInfo},
}

package outer_http

import (
	"myweb/interfaces/http"
	router_tool "myweb/tools"
)

var OuterUrlList = []router_tool.RouterStruct{
	{Method: http.Method.POST, Path: "/user/login", Handler: ApiUserLogin},
	{Method: http.Method.GET, Path: "/user/check_token", Handler: ApiCheckToken},
	{Method: http.Method.POST, Path: "/user/info", Handler: ApiGetUserInfo},
}

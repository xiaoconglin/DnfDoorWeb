package outer_http

import (
	router_tool "myweb/tools"
)

var OuterUrlList = []router_tool.RouterStruct{
	{Method: "POST", Path: "/user/login", Handler: ApiUserLogin},
	{Method: "POST", Path: "/user/info", Handler: ApiGetUserInfo},
}

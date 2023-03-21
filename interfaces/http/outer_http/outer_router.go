package outer_http

import (
	"myweb/interfaces/http"
)

var OuterUrlList = []http.RouterStruct{
	{Method: "POST", Path: "/user/info", Handler: ApiGetUserInfo},
}

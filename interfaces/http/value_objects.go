package http

type method struct {
	GET  string
	POST string
}

var Method = method{
	GET:  "GET",
	POST: "POST",
}

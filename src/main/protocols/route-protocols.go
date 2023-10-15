package main_protocols

import "net/http"

type Route struct {
	Url    string
	Method string
	Func   func(res http.ResponseWriter, req *http.Request)
}

package main_protocols

import (
	"net/http"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
)

type Route struct {
	Url    string
	Method string
	Func   func(w http.ResponseWriter, r *http.Request)
}

type RouteAdapter interface {
	Adapt(controller controller_protocols.Controller) func(w http.ResponseWriter, r *http.Request)
}

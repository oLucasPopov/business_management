package mux_route_adapter

import (
	"fmt"
	"net/http"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	main_protocols "pontos_funcionario/src/main/protocols"
)

type MuxRoute struct {
	main_protocols.RouteAdapter
}

func (mr *MuxRoute) Adapt(controller controller_protocols.Controller) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		res := controller.Handle(&controller_protocols.ControllerRequest{
			Query: r.URL.Query(),
			Body:  r.Body,
		})

		w.WriteHeader(res.StatusCode)
		w.Header().Set("Content-Type", "application/json")
		if res.Body != nil {
			w.Write([]byte(fmt.Sprintf("%s", res.Body)))
		}
	}
}

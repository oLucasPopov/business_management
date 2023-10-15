package app_routes

import (
	employee_routes "pontos_funcionario/src/main/routes/employees"

	"github.com/gorilla/mux"
)

func MakeRoutes() *mux.Router {
	r := mux.NewRouter()

	for _, route := range employee_routes.MakeEmployeeRoutes() {
		r.HandleFunc(route.Url, route.Func).Methods(route.Method)
	}

	return r
}

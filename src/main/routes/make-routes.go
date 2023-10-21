package app_routes

import (
	timekeeping_routes "pontos_funcionario/src/main/routes/employee-timekeeping"
	employee_routes "pontos_funcionario/src/main/routes/employees"

	"github.com/gorilla/mux"
)

func MakeRoutes() *mux.Router {
	r := mux.NewRouter()

	routes := append(
		employee_routes.MakeEmployeeRoutes(),
		timekeeping_routes.MakeTimekeepingRoutes()...,
	)

	for _, route := range routes {
		r.HandleFunc(route.Url, route.Func).Methods(route.Method)
	}

	return r
}

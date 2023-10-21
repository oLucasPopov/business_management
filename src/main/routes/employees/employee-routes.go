package employee_routes

import (
	"net/http"
	main_protocols "pontos_funcionario/src/main/protocols"
)

func MakeEmployeeRoutes() []*main_protocols.Route {
	var routes []*main_protocols.Route
	const employeesUrl string = "/employees"

	routes = append(routes,
		&main_protocols.Route{
			Url:    "/employee/{id}",
			Method: http.MethodGet,
			Func:   GetEmployeeAdapter,
		},
		&main_protocols.Route{
			Url:    employeesUrl,
			Method: http.MethodPost,
			Func:   AddEmployeeAdapter,
		},
		&main_protocols.Route{
			Url:    employeesUrl,
			Method: http.MethodPut,
			Func:   UpdateEmployeeAdapter,
		},
		&main_protocols.Route{
			Url:    employeesUrl,
			Method: http.MethodGet,
			Func:   ListEmployeesAdapter,
		},
	)

	return routes
}

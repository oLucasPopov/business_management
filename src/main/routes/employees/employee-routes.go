package employee_routes

import (
	"net/http"
	employee_factory "pontos_funcionario/src/main/factories"
	main_protocols "pontos_funcionario/src/main/protocols"
	mux_route_adapter "pontos_funcionario/src/main/routes/adapters/mux"
)

func MakeEmployeeRoutes() []*main_protocols.Route {
	var routes []*main_protocols.Route
	const employeesUrl string = "/employees"
	muxRoute := mux_route_adapter.MuxRoute{}

	routes = append(routes,
		&main_protocols.Route{
			Url:    "/employee/{id}",
			Method: http.MethodGet,
			Func:   muxRoute.Adapt(*employee_factory.MakeGetEmployee()),
		},
		&main_protocols.Route{
			Url:    "/employee/{id}",
			Method: http.MethodDelete,
			Func:   muxRoute.Adapt(*employee_factory.MakeDeleteEmployee()),
		},
		&main_protocols.Route{
			Url:    employeesUrl,
			Method: http.MethodPost,
			Func:   muxRoute.Adapt(*employee_factory.MakeAddEmployee()),
		},
		&main_protocols.Route{
			Url:    employeesUrl,
			Method: http.MethodPut,
			Func:   muxRoute.Adapt(*employee_factory.MakeUpdateEmployee()),
		},
		&main_protocols.Route{
			Url:    employeesUrl,
			Method: http.MethodGet,
			Func:   muxRoute.Adapt(*employee_factory.MakeListEmployees()),
		},
	)

	return routes
}

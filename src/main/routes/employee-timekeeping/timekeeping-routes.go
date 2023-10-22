package timekeeping_routes

import (
	"net/http"
	employee_factory "pontos_funcionario/src/main/factories"
	main_protocols "pontos_funcionario/src/main/protocols"
	mux_route_adapter "pontos_funcionario/src/main/routes/adapters/mux"
)

func MakeTimekeepingRoutes() []*main_protocols.Route {
	var routes []*main_protocols.Route
	muxRoute := mux_route_adapter.MuxRoute{}

	routes = append(routes,
		&main_protocols.Route{
			Url:    "/employee/clock-in",
			Method: http.MethodPost,
			Func:   muxRoute.Adapt(*employee_factory.MakeClockIn()),
		},
		&main_protocols.Route{
			Url:    "/employee/clock-out",
			Method: http.MethodPost,
			Func:   muxRoute.Adapt(*employee_factory.MakeClockOut()),
		},
		&main_protocols.Route{
			Url:    "/employee/delete-timekeeping/{id}",
			Method: http.MethodDelete,
			Func:   muxRoute.Adapt(*employee_factory.MakeDeleteTimekeeping()),
		},
		&main_protocols.Route{
			Url:    "/employees/timekeepings",
			Method: http.MethodGet,
			Func:   muxRoute.Adapt(*employee_factory.MakeListTimekeeping()),
		},
	)

	return routes
}

package timekeeping_routes

import (
	"net/http"
	main_protocols "pontos_funcionario/src/main/protocols"
)

func MakeTimekeepingRoutes() []*main_protocols.Route {
	var routes []*main_protocols.Route

	routes = append(routes,
		&main_protocols.Route{
			Url:    "/employee/clock-in",
			Method: http.MethodPost,
			Func:   ClockInRouteAdapter,
		},
		&main_protocols.Route{
			Url:    "/employee/clock-out",
			Method: http.MethodPost,
			Func:   ClockOutRouteAdapter,
		},
		&main_protocols.Route{
			Url:    "/employee/delete-timekeeping/{id}",
			Method: http.MethodDelete,
			Func:   DeleteTimeKeepingRouteAdapter,
		},
		&main_protocols.Route{
			Url:    "/employees/timekeepings",
			Method: http.MethodGet,
			Func:   ListTimekeepingAdapter,
		},
	)

	return routes
}

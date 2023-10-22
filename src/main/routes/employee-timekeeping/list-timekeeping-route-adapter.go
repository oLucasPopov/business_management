package timekeeping_routes

import (
	"fmt"
	"net/http"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	employee_factory "pontos_funcionario/src/main/factories"
)

func ListTimekeepingAdapter(w http.ResponseWriter, r *http.Request) {

	listTimekeeping := employee_factory.MakeListTimekeeping()
	res := listTimekeeping.Handle(&controller_protocols.ControllerRequest{
		Query: r.URL.Query(),
	})

	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	if res.Body != nil {
		w.Write([]byte(fmt.Sprintf("%s", res.Body)))
	}
}

package timekeeping_routes

import (
	"fmt"
	"io"
	"net/http"
	employee_factory "pontos_funcionario/src/main/factories"
)

func ClockInRouteAdapter(w http.ResponseWriter, r *http.Request) {
	reqJson, _ := io.ReadAll(r.Body)
	clockinEmployee := employee_factory.MakeClockIn()

	res := clockinEmployee.Handle(string(reqJson))

	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	if res.Body != nil {
		w.Write([]byte(fmt.Sprintf("%s", res.Body)))
	}
}

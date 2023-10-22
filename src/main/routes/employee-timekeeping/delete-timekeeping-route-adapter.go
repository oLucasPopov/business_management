package timekeeping_routes

import (
	"fmt"
	"net/http"
	employee_factory "pontos_funcionario/src/main/factories"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteTimeKeepingRouteAdapter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	deleteTimekeeping := employee_factory.MakeDeleteTimekeeping()
	res := deleteTimekeeping.Handle(id)
	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	if res.Body != nil {
		w.Write([]byte(fmt.Sprintf("%s", res.Body)))
	}
}

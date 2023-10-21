package employee_routes

import (
	"fmt"
	"net/http"
	employee_factory "pontos_funcionario/src/main/factories"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteEmployeeAdapter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	deleteEmployees := employee_factory.MakeDeleteEmployee()
	res := deleteEmployees.Handle(id)
	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	if res.Body != nil {
		w.Write([]byte(fmt.Sprintf("%s", res.Body)))
	}
}

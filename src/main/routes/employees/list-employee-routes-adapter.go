package employee_routes

import (
	"fmt"
	"net/http"
	employee_factory "pontos_funcionario/src/main/factories"
	"strconv"
)

func ListEmployeesAdapter(w http.ResponseWriter, r *http.Request) {

	pageQuery := r.URL.Query().Get("page")
	page, err := strconv.ParseInt(pageQuery, 10, 32)
	if err != nil {
		page = 1
	}

	listEmployees := employee_factory.MakeListEmployees()
	res := listEmployees.Handle(int32(page))

	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	if res.Body != nil {
		w.Write([]byte(fmt.Sprintf("%s", res.Body)))
	}
}

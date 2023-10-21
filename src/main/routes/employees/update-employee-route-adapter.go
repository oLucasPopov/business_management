package employee_routes

import (
	"fmt"
	"io"
	"net/http"
	employee_factory "pontos_funcionario/src/main/factories"
)

func UpdateEmployeeAdapter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL: ", r.URL, "Method: ", r.Method)

	reqJson, _ := io.ReadAll(r.Body)

	updateEmployee := employee_factory.MakeUpdateEmployee()
	res := updateEmployee.Handle(string(reqJson))

	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	if res.Body != nil {
		w.Write([]byte(fmt.Sprintf("%s", res.Body)))
	}
}

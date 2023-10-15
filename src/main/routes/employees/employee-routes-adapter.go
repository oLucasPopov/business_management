package employee_routes

import (
	"fmt"
	"io"
	"net/http"
	employee_factory "pontos_funcionario/src/main/factories"
	"strconv"

	"github.com/gorilla/mux"
)

func GetEmployeeAdapter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	getEmployees := employee_factory.MakeGetEmployee()
	res := getEmployees.Handle(id)
	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")

	if res.Body != nil {
		w.Write([]byte(fmt.Sprintf("%s", res.Body)))
	}
}

func AddEmployeeAdapter(w http.ResponseWriter, r *http.Request) {
	reqJson, err := io.ReadAll(r.Body)
	fmt.Print(err, reqJson)

	addEmployee := employee_factory.MakeAddEmployee()
	res := addEmployee.Handle(string(reqJson))

	w.WriteHeader(res.StatusCode)
	w.Header().Set("Content-Type", "application/json")
	if res.Body != nil {
		w.Write([]byte(fmt.Sprintf("%s", res.Body)))
	}
}

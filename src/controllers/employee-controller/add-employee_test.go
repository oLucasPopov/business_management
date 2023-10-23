package employee_controller_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	employee_factory "pontos_funcionario/src/main/factories"
	"strings"
	"testing"
)

func TestAddEmployeeFailedValidations(t *testing.T) {
	const employeesRoute string = "/employees"
	const failStatus = "Expected a status code of %d and got %d"

	t.Run("Tests name validation", func(t *testing.T) {
		sut := *employee_factory.MakeAddEmployee()
		req := httptest.NewRequest(http.MethodPost, employeesRoute, strings.NewReader("{}"))

		cr := controller_protocols.ControllerRequest{
			Body: req.Body,
		}
		result := sut.Handle(&cr)

		if result.StatusCode != http.StatusBadRequest {
			fmt.Printf(failStatus, http.StatusBadRequest, result.StatusCode)
			t.Fail()
		}
	})
	t.Run("Tests Salary Type validation", func(t *testing.T) {
		sut := *employee_factory.MakeAddEmployee()
		req := httptest.NewRequest(http.MethodPost, employeesRoute, strings.NewReader(`{"name": "any_name", "salary":1.99}`))

		cr := controller_protocols.ControllerRequest{
			Body: req.Body,
		}
		result := sut.Handle(&cr)

		if result.StatusCode != http.StatusBadRequest {
			fmt.Printf(failStatus, http.StatusBadRequest, result.StatusCode)
			t.Fail()
		}
	})
	t.Run("Tests Invalid Salary Type validation", func(t *testing.T) {
		sut := *employee_factory.MakeAddEmployee()
		req := httptest.NewRequest(http.MethodPost, employeesRoute, strings.NewReader(`{"name": "any_name", "salary_type": "Invalid"}`))

		cr := controller_protocols.ControllerRequest{
			Body: req.Body,
		}
		result := sut.Handle(&cr)

		if result.StatusCode != http.StatusBadRequest {
			fmt.Printf(failStatus, http.StatusBadRequest, result.StatusCode)
			t.Fail()
		}

	})

	t.Run("Tests Salary validation", func(t *testing.T) {
		sut := *employee_factory.MakeAddEmployee()
		req := httptest.NewRequest(http.MethodPost, employeesRoute, strings.NewReader(`{
			"name": "any_name",
			"salary_type": "H"
		}`))

		cr := controller_protocols.ControllerRequest{
			Body: req.Body,
		}
		result := sut.Handle(&cr)

		if result.StatusCode != http.StatusBadRequest {
			fmt.Printf(failStatus, http.StatusBadRequest, result.StatusCode)
			t.Fail()
		}
	})
	t.Run("Tests Salary validation", func(t *testing.T) {

		sut := *employee_factory.MakeAddEmployee()
		req := httptest.NewRequest(http.MethodPost, employeesRoute, strings.NewReader(`{
			"name": "any_name",
			"salary_type": "H",
			"salary": -1
		}`))

		cr := controller_protocols.ControllerRequest{
			Body: req.Body,
		}
		result := sut.Handle(&cr)

		if result.StatusCode != http.StatusBadRequest {
			fmt.Printf(failStatus, http.StatusBadRequest, result.StatusCode)
			t.Fail()
		}
	})

	t.Run("Tests Should not return error if all required fields are informed correctly", func(t *testing.T) {
		sut := *employee_factory.MakeAddEmployee()
		req := httptest.NewRequest(http.MethodPost, "/employees", strings.NewReader(`{
			"name": "any_name",
			"salary_type": "H",
			"salary": 10.33
		}`))

		cr := controller_protocols.ControllerRequest{
			Body: req.Body,
		}
		result := sut.Handle(&cr)

		if result.StatusCode != http.StatusCreated {
			fmt.Printf(failStatus, http.StatusCreated, result.StatusCode)
			t.Fail()
		}
	})
}

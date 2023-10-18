package employee_controller_test

import (
	"fmt"
	"net/http"
	employee_factory "pontos_funcionario/src/main/factories"
	"testing"
)

func TestAddEmployeeFailedValidations(t *testing.T) {
	const failStatus = "Expected a status code of %d and got %d"
	t.Run("Tests name validation", func(t *testing.T) {
		sut := employee_factory.MakeAddEmployee()
		result := sut.Handle("{}")

		if result.StatusCode != http.StatusBadRequest {
			fmt.Printf(failStatus, http.StatusBadRequest, result.StatusCode)
			t.Fail()
		}
	})
	t.Run("Tests Salary Type validation", func(t *testing.T) {
		sut := employee_factory.MakeAddEmployee()
		result := sut.Handle(`{"name": "any_name", "salary":1.99}`)

		if result.StatusCode != http.StatusBadRequest {
			fmt.Printf(failStatus, http.StatusBadRequest, result.StatusCode)
			t.Fail()
		}
	})
	t.Run("Tests Invalid Salary Type validation", func(t *testing.T) {
		sut := employee_factory.MakeAddEmployee()
		result := sut.Handle(`{"name": "any_name", "salary_type": "Invalid"}`)

		if result.StatusCode != http.StatusBadRequest {
			fmt.Printf(failStatus, http.StatusBadRequest, result.StatusCode)
			t.Fail()
		}

	})

	t.Run("Tests Salary validation", func(t *testing.T) {
		sut := employee_factory.MakeAddEmployee()
		result := sut.Handle(`{
			"name": "any_name",
			"salary_type": "H"
		}`)

		if result.StatusCode != http.StatusBadRequest {
			fmt.Printf(failStatus, http.StatusBadRequest, result.StatusCode)
			t.Fail()
		}
	})
	t.Run("Tests Salary validation", func(t *testing.T) {
		sut := employee_factory.MakeAddEmployee()
		result := sut.Handle(`{
			"name": "any_name",
			"salary_type": "H",
			"salary": -1
		}`)

		if result.StatusCode != http.StatusBadRequest {
			fmt.Printf(failStatus, http.StatusBadRequest, result.StatusCode)
			t.Fail()
		}
	})

	t.Run("Tests Should not return error if all required fields are informed correctly", func(t *testing.T) {
		sut := employee_factory.MakeAddEmployee()
		result := sut.Handle(`{
			"name": "any_name",
			"salary_type": "H",
			"salary": 10.33
		}`)

		if result.StatusCode != http.StatusCreated {
			fmt.Printf(failStatus, http.StatusCreated, result.StatusCode)
			t.Fail()
		}
	})
}

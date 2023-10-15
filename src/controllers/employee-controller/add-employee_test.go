package employee_controller_test

import (
	"fmt"
	"net/http"
	employee_factory "pontos_funcionario/src/main/factories"
	"testing"
)

func TestAddEmployeeUnmarshal(t *testing.T) {
	t.Run("Tests unmarshal error", func(t *testing.T) {
		sut := employee_factory.MakeAddEmployee()
		response := sut.Handle("any string")

		if response.StatusCode != http.StatusInternalServerError {
			fmt.Printf("Expected to return status code %d and got %d", http.StatusInternalServerError, response.StatusCode)
			t.Fail()
		}

		err := response.Body
		if err == nil {
			fmt.Printf("Expected to return error but got nil")
			t.Fail()
		}
	})
}

func TestAddEmployeeFailedValidations(t *testing.T) {
	const failStatus = "Expected a status code of %d and got %d"
	const failMessage = "expected %s but got %s"
	const requiredFieldMessage = `the field "%s" is required`
	t.Run("Tests name validation", func(t *testing.T) {
		sut := employee_factory.MakeAddEmployee()
		result := sut.Handle("{}")

		if result.StatusCode != http.StatusBadRequest {
			fmt.Printf(failStatus, http.StatusBadRequest, result.StatusCode)
			t.Fail()
		}

		err := result.Body.(error)

		if err == nil {
			fmt.Printf("Expected to return error but got nil")
			t.Fail()
		}

		requiredFieldValidated := fmt.Sprintf(requiredFieldMessage, "name")
		if err.Error() != requiredFieldValidated {
			fmt.Printf(failMessage, requiredFieldValidated, err.Error())
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

		err := result.Body.(error)

		if err == nil {
			fmt.Printf("Expected to return error but got nil")
			t.Fail()
		}

		requiredFieldValidated := fmt.Sprintf(requiredFieldMessage, "salary_type")
		if err.Error() != requiredFieldValidated {
			fmt.Printf(failMessage, requiredFieldValidated, err.Error())
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

		err := result.Body.(error)

		if err == nil {
			fmt.Printf("Expected to return error but got nil")
			t.Fail()
		}

		requiredFieldValidated := fmt.Sprintf(requiredFieldMessage, "salary")

		if err.Error() != requiredFieldValidated {
			fmt.Printf(failMessage, requiredFieldValidated, err.Error())
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

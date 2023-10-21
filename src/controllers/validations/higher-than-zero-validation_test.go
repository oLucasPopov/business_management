package validations_test

import (
	"fmt"
	"pontos_funcionario/src/controllers/validations"
	"testing"
)

func TestHigherThanZeroValidation(t *testing.T) {
	t.Run("Tests valid Validation", func(t *testing.T) {
		sut := validations.NewHigherThanZeroValidation("salary")
		_, result := sut.Validate(map[string]interface{}{"salary": 1.0})

		if result != nil {
			fmt.Printf("Expected to return nil but got error: %s", result.Error())
			t.Fail()
		}
	})

	t.Run("Tests negative number Validation", func(t *testing.T) {
		sut := validations.NewHigherThanZeroValidation("salary")
		_, result := sut.Validate(map[string]interface{}{"salary": -1})
		expectedError := "the field \"salary\" must be higher than zero"

		if result.Error() != expectedError {
			fmt.Printf("Expected to return %s but got error: %s", expectedError, result.Error())
			t.Fail()
		}
	})
	t.Run("Tests zero Validation", func(t *testing.T) {
		sut := validations.NewHigherThanZeroValidation("salary")
		_, result := sut.Validate(map[string]interface{}{"salary": 0})
		expectedError := "the field \"salary\" must be higher than zero"

		if result.Error() != expectedError {
			fmt.Printf("Expected to return %s but got error: %s", expectedError, result.Error())
			t.Fail()
		}
	})

	t.Run("Tests nil Validation", func(t *testing.T) {
		sut := validations.NewHigherThanZeroValidation("salary")
		_, result := sut.Validate(map[string]interface{}{"not_the_salary": 0})

		if result != nil {
			fmt.Printf("Expected not to return error but got %s", result.Error())
			t.Fail()
		}
	})
}

package validations_test

import (
	"fmt"
	"pontos_funcionario/src/controllers/validations"
	"testing"
)

func TestRequiredValuesValidation(t *testing.T) {
	t.Run("Tests valid values", func(t *testing.T) {
		sut := validations.NewRequiredValuesValidation("salary_type", []interface{}{"H", "M"})
		_, result := sut.Validate("{\"name\": \"any_name\", \"salary_type\": \"H\", \"salary\": 1.99}")

		if result != nil {
			fmt.Printf("Expected to return nil but got error: %s", result.Error())
			t.Fail()
		}
	})
	t.Run("Tests invalid validation", func(t *testing.T) {
		sut := validations.NewRequiredValuesValidation("salary_type", []interface{}{"H", "M"})
		_, result := sut.Validate("{\"name\": \"any_name\", \"salary_type\": \"Z\", \"salary\": 1.99}")

		if result == nil {
			fmt.Printf("Expected to return error but got nil")
			t.Fail()
		}
	})
}

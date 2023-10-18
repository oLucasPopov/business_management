package validations_test

import (
	"fmt"
	"pontos_funcionario/src/controllers/validations"
	"testing"
)

func TestRequiredValuesValidation(t *testing.T) {
	t.Run("Tests valid values", func(t *testing.T) {
		sut := validations.NewRequiredValuesValidation("salary_type", []interface{}{"H", "M"})
		_, result := sut.Validate(
			map[string]interface{}{
				"name":        "any_name",
				"salary":      1.99,
				"salary_type": "H",
			},
		)

		if result != nil {
			fmt.Printf("Expected to return nil but got error: %s", result.Error())
			t.Fail()
		}
	})
	t.Run("Tests invalid validation", func(t *testing.T) {
		sut := validations.NewRequiredValuesValidation("salary_type", []interface{}{"H", "M"})
		_, result := sut.Validate(
			map[string]interface{}{
				"name":        "any_name",
				"salary":      1.99,
				"salary_type": "Z",
			},
		)
		if result == nil {
			fmt.Printf("Expected to return error but got nil")
			t.Fail()
		}
	})
}

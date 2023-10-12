package validations_test

import (
	"fmt"
	"pontos_funcionario/src/controllers/validations"
	"testing"
)

func TestRequiredFieldValidation(t *testing.T) {
	t.Run("Tests valid validation", func(t *testing.T) {
		sut := validations.NewRequiredFieldValidation("name")
		result := sut.Validate("{\"name\": \"any_name\"}")

		if result != nil {
			fmt.Printf("Expected to return nil but got error: %s", result.Error())
			t.Fail()
		}
	})
	t.Run("Tests invalid validation", func(t *testing.T) {
		sut := validations.NewRequiredFieldValidation("name")
		result := sut.Validate("{\"not_the_name\": \"any_name\"}")

		if result == nil {
			fmt.Printf("Expected to return error but got nil")
			t.Fail()
		}
	})
}

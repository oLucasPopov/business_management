package validations_test

import (
	"fmt"
	"pontos_funcionario/src/controllers/validations"
	"testing"
)

func TestEmptyFieldValidation(t *testing.T) {
	t.Run("Tests non empty field Validation", func(t *testing.T) {
		sut := validations.NewEmptyFieldValidation("name")
		_, result := sut.Validate(map[string]interface{}{"name": "any name"})

		if result != nil {
			fmt.Printf("Expected to return nil but got error: %s", result.Error())
			t.Fail()
		}
	})

	t.Run("Tests empty field Validation", func(t *testing.T) {
		sut := validations.NewEmptyFieldValidation("name")
		_, result := sut.Validate(map[string]interface{}{"name": ""})
		expectedError := "the field \"name\" should not be empty"

		if result.Error() != expectedError {
			fmt.Printf("Expected to return %s but got error: %s", expectedError, result.Error())
			t.Fail()
		}
	})

}

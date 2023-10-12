package validations

import (
	"encoding/json"
	"fmt"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
)

type RequiredFieldValidation struct {
	fieldName string
	controller_protocols.Validation
}

func NewRequiredFieldValidation(fieldName string) *RequiredFieldValidation {
	return &RequiredFieldValidation{
		fieldName: fieldName,
	}
}

func (c *RequiredFieldValidation) Validate(input string) error {
	mapInput := map[string]interface{}{}

	_ = json.Unmarshal([]byte(input), &mapInput)

	if mapInput[c.fieldName] == nil {
		return fmt.Errorf(`the field "%s" is required`, c.fieldName)
	}

	return nil
}

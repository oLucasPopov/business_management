package validations

import (
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

func (c *RequiredFieldValidation) Validate(input interface{}) (*string, error) {
	inputMap, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("input is not a map[string]interface{}")
	}

	if inputMap[c.fieldName] == nil {
		return &c.fieldName, fmt.Errorf(`the field "%s" is required`, c.fieldName)
	}

	return nil, nil
}

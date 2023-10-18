package validations

import (
	"fmt"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"slices"
)

type RequiredValuesValidation struct {
	validValues []interface{}
	fieldName   string
	controller_protocols.Validation
}

func NewRequiredValuesValidation(field string, valuesToValidate []interface{}) *RequiredValuesValidation {
	return &RequiredValuesValidation{
		validValues: valuesToValidate,
		fieldName:   field,
	}
}

func (c *RequiredValuesValidation) Validate(input interface{}) (*string, error) {
	inputMap, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("input is not a map[string]interface{}")
	}

	if !slices.Contains(c.validValues, inputMap[c.fieldName]) {
		return &c.fieldName, fmt.Errorf("the value provided for the field %s is invalid", c.fieldName)
	}

	return nil, nil
}

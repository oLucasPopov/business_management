package validations

import (
	"encoding/json"
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

func (c *RequiredValuesValidation) Validate(input string) (*string, error) {
	mapInput := map[string]interface{}{}
	_ = json.Unmarshal([]byte(input), &mapInput)

	if !slices.Contains(c.validValues, mapInput[c.fieldName]) {
		return &c.fieldName, fmt.Errorf("the value provided for the field %s is invalid", c.fieldName)
	}

	return nil, nil
}

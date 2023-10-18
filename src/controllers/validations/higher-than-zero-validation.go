package validations

import (
	"fmt"
)

type HigherThanZeroValidation struct {
	fieldName string
}

func NewHigherThanZeroValidation(fieldName string) *HigherThanZeroValidation {
	return &HigherThanZeroValidation{
		fieldName: fieldName,
	}
}

func (c *HigherThanZeroValidation) Validate(input interface{}) (*string, error) {
	inputMap, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("input is not a map[string]interface{}")
	}

	inputValue, ok := inputMap[c.fieldName].(float64)
	if (!ok) || (inputValue <= 0) {
		return &c.fieldName, fmt.Errorf(`the field "%s" must be higher than zero`, c.fieldName)
	}

	return nil, nil
}

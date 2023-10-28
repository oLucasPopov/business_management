package validations

import "fmt"

type EmptyFieldValidation struct {
	fieldName string
}

func NewEmptyFieldValidation(fieldName string) *EmptyFieldValidation {
	return &EmptyFieldValidation{
		fieldName: fieldName,
	}
}

func (c *EmptyFieldValidation) Validate(input interface{}) (*string, error) {
	inputMap, ok := input.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("input is not a map[string]interface{}")
	}

	if inputMap[c.fieldName] == nil {
		return nil, nil
	}

	inputValue, ok := inputMap[c.fieldName].(string)
	if (!ok) || (inputValue == "") {
		return &c.fieldName, fmt.Errorf(`the field "%s" should not be empty`, c.fieldName)
	}

	return nil, nil
}

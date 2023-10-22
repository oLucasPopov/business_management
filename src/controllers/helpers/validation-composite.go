package controller_helpers

import (
	"encoding/json"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
)

type ValidationComposite struct {
	validations []controller_protocols.Validation
}

func NewValidationComposite(validations []controller_protocols.Validation) *ValidationComposite {
	return &ValidationComposite{
		validations: validations,
	}
}

func (c *ValidationComposite) Validate(input []byte) (*string, error) {
	var inputMap map[string]interface{}
	json.Unmarshal(input, &inputMap)

	for _, v := range c.validations {
		if field, err := v.Validate(inputMap); err != nil {
			return field, err
		}
	}
	return nil, nil
}

package controller_helpers

import controller_protocols "pontos_funcionario/src/controllers/protocols"

type ValidationComposite struct {
	Validations []controller_protocols.Validation
}

func NewValidationComposite(validations []controller_protocols.Validation) *ValidationComposite {
	return &ValidationComposite{
		Validations: validations,
	}
}

func (c *ValidationComposite) Validate(input string) (*string, error) {
	for _, v := range c.Validations {
		if field, err := v.Validate(input); err != nil {
			return field, err
		}
	}
	return nil, nil
}

package employee_factory

import (
	clock_controller "pontos_funcionario/src/controllers/employee-timekeeping-controller"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"pontos_funcionario/src/controllers/validations"
)

func MakeClockIn() *clock_controller.ClockInEmployee {
	requiredFields := []string{"employee_id", "clock_in"}
	validationCollection := []controller_protocols.Validation{}

	for _, required_field := range requiredFields {
		validationCollection = append(validationCollection, validations.NewRequiredFieldValidation(required_field))
	}

	validationCollection = append(validationCollection, validations.NewHigherThanZeroValidation("salary"))

	validationsComposite := controller_helpers.NewValidationComposite(validationCollection)

	return &clock_controller.ClockInEmployee{
		Validations: *validationsComposite,
	}
}

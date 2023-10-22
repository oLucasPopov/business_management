package employee_factory

import (
	clock_controller "pontos_funcionario/src/controllers/employee-timekeeping-controller"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"pontos_funcionario/src/controllers/validations"
	pg_timekeeping_repositories "pontos_funcionario/src/repositories/pg/employee_timekeeping"
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

func MakeClockOut() *controller_protocols.Controller {
	requiredFields := []string{"id", "clock_out"}
	validationCollection := []controller_protocols.Validation{}
	for _, required_field := range requiredFields {
		validationCollection = append(validationCollection, validations.NewRequiredFieldValidation(required_field))
	}
	validationsComposite := controller_helpers.NewValidationComposite(validationCollection)

	clockOutRepository := pg_timekeeping_repositories.ClockOut{}
	closedClockInRepository := pg_timekeeping_repositories.ClosedClockIn{}
	clockOutEmployee := clock_controller.MakeClockOutEmployee(*validationsComposite, clockOutRepository, closedClockInRepository)

	return &clockOutEmployee
}

func MakeDeleteTimekeeping() *controller_protocols.Controller {
	deleteTimeKeepingRepository := pg_timekeeping_repositories.DeleteTimekeeping{}
	deleteTimeKeeping := clock_controller.MakeDeleteTimekeeping(deleteTimeKeepingRepository)
	return &deleteTimeKeeping
}

func MakeListTimekeeping() *controller_protocols.Controller {
	listTimekeepingRepository := pg_timekeeping_repositories.ListTimekeeping{}
	listTimekeeping := clock_controller.MakeListTimekeeping(listTimekeepingRepository)
	return &listTimekeeping
}

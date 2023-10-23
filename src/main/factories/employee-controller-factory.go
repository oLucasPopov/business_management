package employee_factory

import (
	employee_controller "pontos_funcionario/src/controllers/employee-controller"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"pontos_funcionario/src/controllers/validations"
	pg_employee_repositories "pontos_funcionario/src/repositories/pg/employee"
)

func MakeAddEmployee() *controller_protocols.Controller {
	requiredFields := []string{"name", "salary", "salary_type"}
	validationCollection := []controller_protocols.Validation{}

	for _, required_field := range requiredFields {
		validationCollection = append(validationCollection, validations.NewRequiredFieldValidation(required_field))
	}

	validationCollection = append(validationCollection,
		validations.NewRequiredValuesValidation("salary_type", []interface{}{"H", "M"}),
		validations.NewHigherThanZeroValidation("salary"))

	validationsComposite := controller_helpers.NewValidationComposite(validationCollection)
	addEmployeeRepository := pg_employee_repositories.AddEmployee{}
	addEmployee := employee_controller.MakeAddEmployee(addEmployeeRepository, *validationsComposite)
	return &addEmployee
}
func MakeUpdateEmployee() *controller_protocols.Controller {
	requiredFields := []string{"name", "salary", "salary_type"}
	validationCollection := []controller_protocols.Validation{}

	for _, required_field := range requiredFields {
		validationCollection = append(validationCollection, validations.NewRequiredFieldValidation(required_field))
	}

	validationCollection = append(validationCollection,
		validations.NewRequiredValuesValidation("salary_type", []interface{}{"H", "M"}),
		validations.NewHigherThanZeroValidation("salary"))

	validationsComposite := controller_helpers.NewValidationComposite(validationCollection)

	updateEmployeeRepository := pg_employee_repositories.UpdateEmployee{}
	updateEmployee := employee_controller.MakeUpdateEmployee(updateEmployeeRepository, *validationsComposite)
	return &updateEmployee
}

func MakeGetEmployee() employee_controller.GetEmployee {
	return employee_controller.GetEmployee{
		GetEmployeeRepository: pg_employee_repositories.GetEmployee{},
	}
}

func MakeListEmployees() employee_controller.ListEmployees {
	return employee_controller.ListEmployees{
		ListEmployeesRepository: pg_employee_repositories.ListEmployees{},
	}
}

func MakeDeleteEmployee() employee_controller.DeleteEmployee {
	return employee_controller.DeleteEmployee{
		DeleteEmployeeRepository: pg_employee_repositories.DeleteEmployee{},
	}
}

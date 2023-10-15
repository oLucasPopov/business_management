package employee_factory

import (
	employee_controller "pontos_funcionario/src/controllers/employee-controller"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"pontos_funcionario/src/controllers/validations"
	pg_employee_repositories "pontos_funcionario/src/repositories/pg/employee"
)

func MakeAddEmployee() employee_controller.AddEmployee {
	required_fields := []string{"name", "salary", "salary_type"}
	validation_collection := []controller_protocols.Validation{}

	for _, required_field := range required_fields {
		validation_collection = append(validation_collection, validations.NewRequiredFieldValidation(required_field))
	}

	validation_collection = append(validation_collection,
		validations.NewRequiredValuesValidation("salary_type", []interface{}{"H", "M"}))

	validationsComposite := controller_helpers.NewValidationComposite(validation_collection)

	return employee_controller.AddEmployee{
		EmployeeRepository: pg_employee_repositories.AddEmployee{},
		Validations:        *validationsComposite,
	}
}

func MakeGetEmployee() employee_controller.GetEmployee {
	return employee_controller.GetEmployee{
		GetEmployeeRepository: pg_employee_repositories.GetEmployee{},
	}
}

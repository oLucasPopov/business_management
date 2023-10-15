package employee_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"pontos_funcionario/src/models"
	pg_employee_repositories "pontos_funcionario/src/repositories/pg/employee"
)

type AddEmployee struct {
	EmployeeRepository pg_employee_repositories.AddEmployee
	Validations        controller_helpers.ValidationComposite
}

func (c *AddEmployee) Handle(request string) controller_protocols.ControllerResponse {
	addEmployee := models.AddEmployee{}
	err := json.Unmarshal([]byte(request), &addEmployee)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	err = c.Validations.Validate(request)

	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	if *addEmployee.Salary < 0 {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, errors.New("the field salary must be higher than zero"))
	}

	newEmployee, err := c.EmployeeRepository.Add(addEmployee)

	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, errors.New("error while creating employee: "+err.Error()))
	}

	jsonEmployee, err := json.Marshal(newEmployee)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}
	return controller_protocols.ControllerResponse{
		StatusCode: http.StatusCreated,
		Body:       jsonEmployee,
	}
}

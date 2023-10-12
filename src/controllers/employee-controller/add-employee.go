package employee_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"pontos_funcionario/src/models"
	pg_repositories "pontos_funcionario/src/repositories/pg"
)

type AddEmployee struct {
	EmployeeRepository pg_repositories.Employee
	Validation         controller_helpers.ValidationComposite
}

func (c *AddEmployee) Handle(request string) controller_protocols.ControllerResponse {
	addEmployee := models.AddEmployee{}
	err := json.Unmarshal([]byte(request), &addEmployee)
	if err != nil {
		return controller_protocols.ControllerResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       errors.New("coudn't unmarshal json"),
		}
	}

	for _, validation := range c.Validation.Validations {
		if err := validation.Validate(request); err != nil {
			return controller_protocols.ControllerResponse{
				StatusCode: http.StatusBadRequest,
				Body:       err,
			}
		}
	}

	if *addEmployee.Salary < 0 {
		return controller_protocols.ControllerResponse{
			StatusCode: http.StatusBadRequest,
			Body:       errors.New("the field salary must be higher than zero"),
		}
	}

	newEmployee, err := c.EmployeeRepository.Add(addEmployee)

	if err != nil {
		return controller_protocols.ControllerResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       errors.New("error while creating employee: " + err.Error()),
		}
	}

	return controller_protocols.ControllerResponse{
		StatusCode: http.StatusCreated,
		Body:       newEmployee,
	}
}

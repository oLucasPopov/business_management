package employee_controller

import (
	"encoding/json"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"pontos_funcionario/src/models"
	pg_employee_repositories "pontos_funcionario/src/repositories/pg/employee"
)

type UpdateEmployee struct {
	UpdateEmployeeRepository pg_employee_repositories.UpdateEmployee
	Validations              controller_helpers.ValidationComposite
}

func (c *UpdateEmployee) Handle(request string) controller_protocols.ControllerResponse {
	fieldError, err := c.Validations.Validate(request)

	if err != nil {
		return *controller_helpers.ErrorFieldResponse(http.StatusBadRequest, err, *fieldError)
	}

	employee := models.Employee{}
	err = json.Unmarshal([]byte(request), &employee)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	employee, err = c.UpdateEmployeeRepository.Handle(employee)
	if err != nil {
		return controller_protocols.ControllerResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}

	jsonEmployee, err := json.Marshal(employee)

	if err != nil {
		return controller_protocols.ControllerResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}

	return controller_protocols.ControllerResponse{
		StatusCode: http.StatusOK,
		Body:       jsonEmployee,
	}
}

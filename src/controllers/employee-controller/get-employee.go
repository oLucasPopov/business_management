package employee_controller

import (
	"encoding/json"
	"net/http"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	pg_employee_repositories "pontos_funcionario/src/repositories/pg/employee"
)

type GetEmployee struct {
	GetEmployeeRepository pg_employee_repositories.GetEmployee
}

func (c *GetEmployee) Handle(id int64) controller_protocols.ControllerResponse {
	employee, err := c.GetEmployeeRepository.Handle(id)
	if err != nil {
		return controller_protocols.ControllerResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}

	if employee == nil {
		return controller_protocols.ControllerResponse{
			StatusCode: http.StatusNotFound,
			Body:       nil,
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

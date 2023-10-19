package employee_controller

import (
	"encoding/json"
	"net/http"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	pg_employee_repositories "pontos_funcionario/src/repositories/pg/employee"
)

type ListEmployees struct {
	ListEmployeesRepository pg_employee_repositories.ListEmployees
}

func (c *ListEmployees) Handle(page int32) controller_protocols.ControllerResponse {
	employees, err := c.ListEmployeesRepository.Handle(page)
	if err != nil {
		return controller_protocols.ControllerResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
	}

	jsonEmployee, err := json.Marshal(employees)
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

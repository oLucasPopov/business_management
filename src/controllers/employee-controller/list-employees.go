package employee_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	pg_employee_repositories "pontos_funcionario/src/repositories/pg/employee"
	"strconv"
)

type ListEmployees struct {
	listEmployeesRepository pg_employee_repositories.ListEmployees
}

func MakeListEmployees(listEmployeesRepository pg_employee_repositories.ListEmployees) controller_protocols.Controller {
	return &ListEmployees{
		listEmployeesRepository: listEmployeesRepository,
	}
}

func (c *ListEmployees) Handle(request *controller_protocols.ControllerRequest) controller_protocols.ControllerResponse {
	page, err := strconv.ParseInt(request.Query.Get("page"), 10, 32)
	if err != nil {
		return *controller_helpers.ErrorFieldResponse(http.StatusBadRequest, errors.New(`the query "page" must be an integer`), "id")
	}

	employees, err := c.listEmployeesRepository.Handle(int32(page))
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

package employee_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	employee_repositories_protocols "pontos_funcionario/src/repositories/protocols/employees"
	"strconv"
)

type ListEmployees struct {
	listEmployeesRepository employee_repositories_protocols.ListEmployees
}

func MakeListEmployees(listEmployeesRepository employee_repositories_protocols.ListEmployees) controller_protocols.Controller {
	return &ListEmployees{
		listEmployeesRepository: listEmployeesRepository,
	}
}

func (c *ListEmployees) Handle(request *controller_protocols.ControllerRequest) controller_protocols.ControllerResponse {
	pageStr := request.Query.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.ParseInt(pageStr, 10, 32)
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

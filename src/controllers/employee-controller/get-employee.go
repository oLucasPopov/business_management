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

type GetEmployee struct {
	getEmployeeRepository employee_repositories_protocols.GetEmployee
}

func MakeGetEmployee(getEmployeeRepository employee_repositories_protocols.GetEmployee) controller_protocols.Controller {
	return &GetEmployee{
		getEmployeeRepository: getEmployeeRepository,
	}
}

func (c *GetEmployee) Handle(request *controller_protocols.ControllerRequest) controller_protocols.ControllerResponse {
	id, err := strconv.ParseInt(request.Params["id"], 10, 64)
	if err != nil {
		return *controller_helpers.ErrorFieldResponse(http.StatusBadRequest, errors.New(`the param "id" must be an integer`), "id")
	}

	employee, err := c.getEmployeeRepository.Handle(id)
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

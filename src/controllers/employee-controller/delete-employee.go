package employee_controller

import (
	"errors"
	"fmt"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	employee_repositories_protocols "pontos_funcionario/src/repositories/protocols/employees"
	"strconv"
)

type DeleteEmployee struct {
	deleteEmployeeRepository employee_repositories_protocols.DeleteEmployee
	controller_protocols.Controller
}

func MakeDeleteEmployee(deleteEmployeeRepository employee_repositories_protocols.DeleteEmployee) controller_protocols.Controller {
	return &DeleteEmployee{
		deleteEmployeeRepository: deleteEmployeeRepository,
	}
}

func (c *DeleteEmployee) Handle(request *controller_protocols.ControllerRequest) controller_protocols.ControllerResponse {
	id, err := strconv.ParseInt(request.Params["id"], 10, 64)
	if err != nil {
		return *controller_helpers.ErrorFieldResponse(http.StatusBadRequest, errors.New(`the param "id" must be an integer`), "id")
	}

	err = c.deleteEmployeeRepository.Handle(id)

	if err != nil {
		fmt.Println(err)
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	return controller_protocols.ControllerResponse{
		StatusCode: http.StatusOK,
	}
}

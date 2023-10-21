package employee_controller

import (
	"fmt"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	pg_employee_repositories "pontos_funcionario/src/repositories/pg/employee"
)

type DeleteEmployee struct {
	DeleteEmployeeRepository pg_employee_repositories.DeleteEmployee
}

func (c *DeleteEmployee) Handle(id int64) controller_protocols.ControllerResponse {
	err := c.DeleteEmployeeRepository.Handle(id)

	if err != nil {
		fmt.Println(err)
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	return controller_protocols.ControllerResponse{
		StatusCode: http.StatusOK,
	}
}

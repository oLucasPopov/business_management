package clock_controller

import (
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	pg_timekeeping_repositories "pontos_funcionario/src/repositories/pg/employee_timekeeping"
)

type DeleteTimeKeeping struct {
	DeleteTimeKeepingRepository pg_timekeeping_repositories.DeleteTimekeeping
}

func (c *DeleteTimeKeeping) Handle(id int64) controller_protocols.ControllerResponse {
	err := c.DeleteTimeKeepingRepository.Handle(id)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	return controller_protocols.ControllerResponse{
		StatusCode: http.StatusOK,
	}
}

package clock_controller

import (
	"errors"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	pg_timekeeping_repositories "pontos_funcionario/src/repositories/pg/employee_timekeeping"
	"strconv"
)

type DeleteTimeKeeping struct {
	deleteTimeKeepingRepository pg_timekeeping_repositories.DeleteTimekeeping
	controller_protocols.Controller
}

func MakeDeleteTimekeeping(deleteTimeKeepingRepository pg_timekeeping_repositories.DeleteTimekeeping) controller_protocols.Controller {
	return &DeleteTimeKeeping{
		deleteTimeKeepingRepository: deleteTimeKeepingRepository,
	}
}

func (c *DeleteTimeKeeping) Handle(request *controller_protocols.ControllerRequest) controller_protocols.ControllerResponse {
	id, err := strconv.ParseInt(request.Params["id"], 10, 64)
	if err != nil {
		return *controller_helpers.ErrorFieldResponse(http.StatusBadRequest, errors.New(`the param "id" must be an integer`), "id")
	}

	err = c.deleteTimeKeepingRepository.Handle(id)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	return controller_protocols.ControllerResponse{
		StatusCode: http.StatusOK,
	}
}

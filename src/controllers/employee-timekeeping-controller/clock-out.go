package clock_controller

import (
	"encoding/json"
	"errors"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"pontos_funcionario/src/models"
	pg_timekeeping_repositories "pontos_funcionario/src/repositories/pg/employee_timekeeping"
)

type ClockOutEmployee struct {
	Validations             controller_helpers.ValidationComposite
	ClockOutRepository      pg_timekeeping_repositories.ClockOut
	ClosedClockInRepository pg_timekeeping_repositories.ClosedClockIn
}

func (c *ClockOutEmployee) Handle(request string) controller_protocols.ControllerResponse {
	field, err := c.Validations.Validate(request)
	if err != nil {
		return *controller_helpers.ErrorFieldResponse(http.StatusBadRequest, err, *field)
	}

	var addClockOut models.AddClockOutEmployee

	err = json.Unmarshal([]byte(request), &addClockOut)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	closed, err := c.ClosedClockInRepository.Handle(*addClockOut.Id)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	if *closed {
		return *controller_helpers.ErrorResponse(http.StatusConflict, errors.New("this timekeeping is already clocked out"))
	}

	timeKeeping, err := c.ClockOutRepository.Handle(addClockOut)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	jsonTimeKeeping, err := json.Marshal(timeKeeping)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	return controller_protocols.ControllerResponse{
		StatusCode: http.StatusOK,
		Body:       jsonTimeKeeping,
	}
}

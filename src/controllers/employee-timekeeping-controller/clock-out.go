package clock_controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"pontos_funcionario/src/models"
	pg_timekeeping_repositories "pontos_funcionario/src/repositories/pg/employee_timekeeping"
)

type ClockOutEmployee struct {
	validations             controller_helpers.ValidationComposite
	clockOutRepository      pg_timekeeping_repositories.ClockOut
	closedClockInRepository pg_timekeeping_repositories.ClosedClockIn
	controller_protocols.Controller
}

func MakeClockOutEmployee(
	validations controller_helpers.ValidationComposite,
	clockOutRepository pg_timekeeping_repositories.ClockOut,
	closedClockInRepository pg_timekeeping_repositories.ClosedClockIn,
) controller_protocols.Controller {
	return &ClockOutEmployee{
		validations:             validations,
		clockOutRepository:      clockOutRepository,
		closedClockInRepository: closedClockInRepository,
	}
}

func (c *ClockOutEmployee) Handle(request *controller_protocols.ControllerRequest) controller_protocols.ControllerResponse {
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	field, err := c.validations.Validate(requestBody)
	if err != nil {
		return *controller_helpers.ErrorFieldResponse(http.StatusBadRequest, err, *field)
	}

	var addClockOut models.AddClockOutEmployee

	err = json.Unmarshal(requestBody, &addClockOut)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	closed, err := c.closedClockInRepository.Handle(*addClockOut.Id)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	if *closed {
		return *controller_helpers.ErrorResponse(http.StatusConflict, errors.New("this timekeeping is already clocked out"))
	}

	timeKeeping, err := c.clockOutRepository.Handle(addClockOut)
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

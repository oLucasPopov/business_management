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

type ClockInEmployee struct {
	clockInRepository     pg_timekeeping_repositories.ClockIn
	openClockInRepository pg_timekeeping_repositories.OpenClockIn
	validations           controller_helpers.ValidationComposite
	controller_protocols.Controller
}

func MakeClockInEmployee(
	clockInRepository pg_timekeeping_repositories.ClockIn,
	openClockInRepository pg_timekeeping_repositories.OpenClockIn,
	validations controller_helpers.ValidationComposite,
) controller_protocols.Controller {
	return &ClockInEmployee{
		clockInRepository:     clockInRepository,
		openClockInRepository: openClockInRepository,
		validations:           validations,
	}
}

func (c *ClockInEmployee) Handle(request *controller_protocols.ControllerRequest) controller_protocols.ControllerResponse {
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	fieldErr, err := c.validations.Validate(requestBody)
	if err != nil {
		return *controller_helpers.ErrorFieldResponse(http.StatusBadRequest, err, *fieldErr)
	}

	var clockIn = models.AddClockInEmployee{}
	err = json.Unmarshal(requestBody, &clockIn)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	openClockIns, err := c.openClockInRepository.Handle(*clockIn.EmployeeId)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	if *openClockIns > 0 {
		return *controller_helpers.ErrorResponse(http.StatusConflict, errors.New("this employee already has an open clock-in"))
	}

	timekeeping, err := c.clockInRepository.Handle(clockIn)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	jsonTimeKeeping, err := json.Marshal(timekeeping)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	return controller_protocols.ControllerResponse{
		StatusCode: http.StatusCreated,
		Body:       jsonTimeKeeping,
	}
}

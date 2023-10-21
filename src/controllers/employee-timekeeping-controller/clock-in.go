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

type ClockInEmployee struct {
	ClockInRepository     pg_timekeeping_repositories.ClockIn
	OpenClockInRepository pg_timekeeping_repositories.OpenClockIn
	Validations           controller_helpers.ValidationComposite
}

func (c *ClockInEmployee) Handle(request string) controller_protocols.ControllerResponse {
	fieldErr, err := c.Validations.Validate(request)
	if err != nil {
		return *controller_helpers.ErrorFieldResponse(http.StatusBadRequest, err, *fieldErr)
	}

	var clockIn = models.AddClockInEmployee{}
	err = json.Unmarshal([]byte(request), &clockIn)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	openClockIns, err := c.OpenClockInRepository.Handle(*clockIn.EmployeeId)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	if *openClockIns > 0 {
		return *controller_helpers.ErrorResponse(http.StatusConflict, errors.New("this employee already has an open clock-in"))
	}

	timekeeping, err := c.ClockInRepository.Handle(clockIn)
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

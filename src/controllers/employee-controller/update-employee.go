package employee_controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"pontos_funcionario/src/models"
	employee_repositories_protocols "pontos_funcionario/src/repositories/protocols/employees"
)

type UpdateEmployee struct {
	updateEmployeeRepository employee_repositories_protocols.UpdateEmployee
	getEmployeeRepository    employee_repositories_protocols.GetEmployee
	validations              controller_helpers.ValidationComposite
}

func MakeUpdateEmployee(updateEmployeeRepository employee_repositories_protocols.UpdateEmployee,
	getEmployeeRepository employee_repositories_protocols.GetEmployee,
	validations controller_helpers.ValidationComposite) controller_protocols.Controller {
	return &UpdateEmployee{
		updateEmployeeRepository: updateEmployeeRepository,
		getEmployeeRepository:    getEmployeeRepository,
		validations:              validations,
	}
}

func (c *UpdateEmployee) Handle(request *controller_protocols.ControllerRequest) controller_protocols.ControllerResponse {
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	fieldErr, err := c.validations.Validate(requestBody)
	if err != nil {
		return *controller_helpers.ErrorFieldResponse(http.StatusBadRequest, err, *fieldErr)
	}

	employee := models.Employee{}
	err = json.Unmarshal(requestBody, &employee)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}
	existingEmployee, err := c.getEmployeeRepository.Handle(employee.Id)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}

	if existingEmployee == nil {
		return *controller_helpers.ErrorResponse(http.StatusNotFound, fmt.Errorf("the employee %d does not exist", employee.Id))
	}

	employee, err = c.updateEmployeeRepository.Handle(employee)
	if err != nil {
		return controller_protocols.ControllerResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
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

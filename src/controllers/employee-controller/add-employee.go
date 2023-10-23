package employee_controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	controller_helpers "pontos_funcionario/src/controllers/helpers"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
	"pontos_funcionario/src/models"
	employee_repositories_protocols "pontos_funcionario/src/repositories/protocols/employees"
)

type AddEmployee struct {
	employeeRepository employee_repositories_protocols.AddEmployee
	validations        controller_helpers.ValidationComposite
	controller_protocols.Controller
}

func MakeAddEmployee(employeeRepository employee_repositories_protocols.AddEmployee,
	validations controller_helpers.ValidationComposite) controller_protocols.Controller {
	return &AddEmployee{
		employeeRepository: employeeRepository,
		validations:        validations,
	}
}

func (c *AddEmployee) Handle(request *controller_protocols.ControllerRequest) controller_protocols.ControllerResponse {
	reqJson, err := io.ReadAll(request.Body)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	fieldError, err := c.validations.Validate(reqJson)
	if err != nil {
		return *controller_helpers.ErrorFieldResponse(http.StatusBadRequest, err, *fieldError)
	}

	addEmployee := models.AddEmployee{}
	err = json.Unmarshal(reqJson, &addEmployee)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusBadRequest, err)
	}

	newEmployee, err := c.employeeRepository.Add(addEmployee)

	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, errors.New("error while creating employee: "+err.Error()))
	}

	jsonEmployee, err := json.Marshal(newEmployee)
	if err != nil {
		return *controller_helpers.ErrorResponse(http.StatusInternalServerError, err)
	}
	return controller_protocols.ControllerResponse{
		StatusCode: http.StatusCreated,
		Body:       jsonEmployee,
	}
}

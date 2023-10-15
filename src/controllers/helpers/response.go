package controller_helpers

import (
	"encoding/json"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
)

type errorResponse struct {
	Error string `json:"error"`
	Field string `json:"field,omitempty"`
}

func controllerResponse(statusCode int, response []byte) *controller_protocols.ControllerResponse {
	return &controller_protocols.ControllerResponse{
		StatusCode: int(statusCode),
		Body:       response,
	}
}

func ErrorResponse(statusCode int64, err error) *controller_protocols.ControllerResponse {
	jsonRes, _ := json.Marshal(errorResponse{
		Error: err.Error(),
	})

	return controllerResponse(int(statusCode), jsonRes)
}
func ErrorFieldResponse(statusCode int64, err error, field string) *controller_protocols.ControllerResponse {
	jsonRes, _ := json.Marshal(errorResponse{
		Error: err.Error(),
		Field: field,
	})

	return controllerResponse(int(statusCode), jsonRes)
}

package controller_helpers

import (
	"encoding/json"
	controller_protocols "pontos_funcionario/src/controllers/protocols"
)

type errorResponse struct {
	Error string `json:"error"`
	Field string `json:"field,omitempty"`
}

func ErrorResponse(statusCode int64, err error) *controller_protocols.ControllerResponse {
	jsonRes, _ := json.Marshal(errorResponse{
		Error: err.Error(),
	})

	return &controller_protocols.ControllerResponse{
		StatusCode: int(statusCode),
		Body:       jsonRes,
	}
}

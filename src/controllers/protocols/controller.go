package controller_protocols

type Controller interface {
	Handle(*ControllerRequest) ControllerResponse
}

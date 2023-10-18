package controller_protocols

type Validation interface {
	Validate(interface{}) (*string, error)
}

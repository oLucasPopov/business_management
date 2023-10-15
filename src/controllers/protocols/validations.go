package controller_protocols

type Validation interface {
	Validate(string) (*string, error)
}

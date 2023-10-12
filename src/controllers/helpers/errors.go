package controller_helpers

import (
	"fmt"
)

func MissingFieldError(field string) error {
	return fmt.Errorf("the field %s is required", field)
}

package controller_protocols

import (
	"io"
	"net/url"
)

type ControllerRequest struct {
	Body   io.ReadCloser
	Query  url.Values
	Params map[string]string
}

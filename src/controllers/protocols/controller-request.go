package controller_protocols

import "net/url"

type ControllerRequest struct {
	Body  interface{}
	Query url.Values
}

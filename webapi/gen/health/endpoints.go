// Code generated by goa v3.7.6, DO NOT EDIT.
//
// health endpoints
//
// Command:
// $ goa gen github.com/is-hoku/goa-sample/webapi/design

package health

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "health" service endpoints.
type Endpoints struct {
	Check goa.Endpoint
}

// NewEndpoints wraps the methods of the "health" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Check: NewCheckEndpoint(s),
	}
}

// Use applies the given middleware to all the "health" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Check = m(e.Check)
}

// NewCheckEndpoint returns an endpoint function that calls the method "check"
// of service "health".
func NewCheckEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		res, err := s.Check(ctx)
		if err != nil {
			return nil, err
		}
		vres := NewViewedHealthResult(res, "default")
		return vres, nil
	}
}

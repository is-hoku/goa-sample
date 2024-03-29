// Code generated by goa v3.7.6, DO NOT EDIT.
//
// health service
//
// Command:
// $ goa gen github.com/is-hoku/goa-sample/webapi/design

package health

import (
	"context"

	healthviews "github.com/is-hoku/goa-sample/webapi/gen/health/views"
)

// Service is the health service interface.
type Service interface {
	// ヘルスチェック
	Check(context.Context) (res *HealthResult, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "health"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"check"}

// CustomError is the error returned error name and message.
type CustomError struct {
	// Name of error
	Name string
	// Message of error
	Message string
}

// HealthResult is the result type of the health service check method.
type HealthResult struct {
	// health message
	Message string
}

// Error returns an error description.
func (e *CustomError) Error() string {
	return "CustomError is the error returned error name and message."
}

// ErrorName returns "CustomError".
func (e *CustomError) ErrorName() string {
	return e.Name
}

// NewHealthResult initializes result type HealthResult from viewed result type
// HealthResult.
func NewHealthResult(vres *healthviews.HealthResult) *HealthResult {
	return newHealthResult(vres.Projected)
}

// NewViewedHealthResult initializes viewed result type HealthResult from
// result type HealthResult using the given view.
func NewViewedHealthResult(res *HealthResult, view string) *healthviews.HealthResult {
	p := newHealthResultView(res)
	return &healthviews.HealthResult{Projected: p, View: "default"}
}

// newHealthResult converts projected type HealthResult to service type
// HealthResult.
func newHealthResult(vres *healthviews.HealthResultView) *HealthResult {
	res := &HealthResult{}
	if vres.Message != nil {
		res.Message = *vres.Message
	}
	return res
}

// newHealthResultView projects result type HealthResult to projected type
// HealthResultView using the "default" view.
func newHealthResultView(res *HealthResult) *healthviews.HealthResultView {
	vres := &healthviews.HealthResultView{
		Message: &res.Message,
	}
	return vres
}

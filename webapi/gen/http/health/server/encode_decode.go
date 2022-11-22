// Code generated by goa v3.7.6, DO NOT EDIT.
//
// health HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/is-hoku/goa-sample/webapi/design

package server

import (
	"context"
	"errors"
	"net/http"

	health "github.com/is-hoku/goa-sample/webapi/gen/health"
	healthviews "github.com/is-hoku/goa-sample/webapi/gen/health/views"
	goahttp "goa.design/goa/v3/http"
)

// EncodeCheckResponse returns an encoder for responses returned by the health
// check endpoint.
func EncodeCheckResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*healthviews.HealthResult)
		enc := encoder(ctx, w)
		body := NewCheckResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeCheckError returns an encoder for errors returned by the check health
// endpoint.
func EncodeCheckError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		var en ErrorNamer
		if !errors.As(v, &en) {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "internal_error":
			var res *health.CustomError
			errors.As(v, &res)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewCheckInternalErrorResponseBody(res)
			}
			w.Header().Set("goa-error", res.ErrorName())
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}
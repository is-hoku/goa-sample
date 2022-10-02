package main

import (
	health "github.com/is-hoku/goa-sample/webapi/gen/health"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

func healthCustomErrorResponse(err error) goahttp.Statuser {
	// Error Handling for Decoding & Validation
	if serr, ok := err.(*goa.ServiceError); ok {
		switch serr.Name {
		case "missing_payload":
			return &health.CustomError{Name: "bad_request", Message: "Missing Payload"}
		case "decode_payload":
			return &health.CustomError{Name: "bad_request", Message: "Invalid Body"}
		case "invalid_field_type":
			return &health.CustomError{Name: "bad_request", Message: "Invalid Field Type"}
		case "missing_field":
			return &health.CustomError{Name: "unauthorized", Message: "Unauthorized"}
		case "invalid_enum_value":
			return &health.CustomError{Name: "bad_request", Message: "Invalid Value of a Payload"}
		case "invalid_format":
			return &health.CustomError{Name: "bad_request", Message: "Invalid Format"}
		case "invalid_pattern":
			return &health.CustomError{Name: "bad_request", Message: "Invalid Value of a Payload"}
		case "invalid_range":
			return &health.CustomError{Name: "bad_request", Message: "Invalid Value of a Payload"}
		case "invalid_length":
			return &health.CustomError{Name: "bad_request", Message: "Invalid Value of a Payload"}
		default:
			return &health.CustomError{Name: "internal_error", Message: "Internal Server Error"}
		}
	} else if serr, ok := err.(*health.CustomError); ok { // Error Handling for Business logic
		return &health.CustomError{Name: serr.Name, Message: serr.Message}
	}
	return &health.CustomError{Name: "internal_error", Message: "Internal Server Error"}
}

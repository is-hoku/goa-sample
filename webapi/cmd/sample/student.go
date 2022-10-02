package main

import (
	student "github.com/is-hoku/goa-sample/webapi/gen/student"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

func studentCustomErrorResponse(err error) goahttp.Statuser {
	// Error Handling for Decoding & Validation
	if serr, ok := err.(*goa.ServiceError); ok {
		switch serr.Name {
		case "missing_payload":
			return &student.CustomError{Name: "bad_request", Message: "Missing Payload"}
		case "decode_payload":
			return &student.CustomError{Name: "bad_request", Message: "Invalid Body"}
		case "invalid_field_type":
			return &student.CustomError{Name: "bad_request", Message: "Invalid Field Type"}
		case "missing_field":
			return &student.CustomError{Name: "unauthorized", Message: "Unauthorized"}
		case "invalid_enum_value":
			return &student.CustomError{Name: "bad_request", Message: "Invalid Value of a Payload"}
		case "invalid_format":
			return &student.CustomError{Name: "bad_request", Message: "Invalid Format"}
		case "invalid_pattern":
			return &student.CustomError{Name: "bad_request", Message: "Invalid Value of a Payload"}
		case "invalid_range":
			return &student.CustomError{Name: "bad_request", Message: "Invalid Value of a Payload"}
		case "invalid_length":
			return &student.CustomError{Name: "bad_request", Message: "Invalid Value of a Payload"}
		default:
			return &student.CustomError{Name: "internal_error", Message: "Internal Server Error"}
		}
	} else if serr, ok := err.(*student.CustomError); ok { // Error Handling for Business logic
		return &student.CustomError{Name: serr.Name, Message: serr.Message}
	}
	return &student.CustomError{Name: "internal_error", Message: "Internal Server Error"}
}

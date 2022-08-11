package student

import "net/http"

func (err CustomError) StatusCode() int {
	switch err.Name {
	case "not_found":
		return http.StatusNotFound
	case "internal_error":
		return http.StatusInternalServerError
	case "bad_request":
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

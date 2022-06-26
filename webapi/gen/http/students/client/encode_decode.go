// Code generated by goa v3.7.6, DO NOT EDIT.
//
// students HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/is-hoku/goa-template/webapi/design

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	students "github.com/is-hoku/goa-template/gen/students"
	studentsviews "github.com/is-hoku/goa-template/gen/students/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildGetStudentRequest instantiates a HTTP request object with method and
// path set to call the "students" service "get student" endpoint
func (c *Client) BuildGetStudentRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int64
	)
	{
		p, ok := v.(*students.GetStudentPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("students", "get student", "*students.GetStudentPayload", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetStudentStudentsPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("students", "get student", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetStudentResponse returns a decoder for responses returned by the
// students get student endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeGetStudentResponse may return the following errors:
//	- "internal_error" (type *students.CustomError): http.StatusInternalServerError
//	- "not_found" (type *students.CustomError): http.StatusNotFound
//	- error: internal error
func DecodeGetStudentResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetStudentResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("students", "get student", err)
			}
			p := NewGetStudentStudentOK(&body)
			view := "default"
			vres := &studentsviews.Student{Projected: p, View: view}
			if err = studentsviews.ValidateStudent(vres); err != nil {
				return nil, goahttp.ErrValidationError("students", "get student", err)
			}
			res := students.NewStudent(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GetStudentInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("students", "get student", err)
			}
			err = ValidateGetStudentInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("students", "get student", err)
			}
			return nil, NewGetStudentInternalError(&body)
		case http.StatusNotFound:
			var (
				body GetStudentNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("students", "get student", err)
			}
			err = ValidateGetStudentNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("students", "get student", err)
			}
			return nil, NewGetStudentNotFound(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("students", "get student", resp.StatusCode, string(body))
		}
	}
}

// BuildGetStudentsRequest instantiates a HTTP request object with method and
// path set to call the "students" service "get students" endpoint
func (c *Client) BuildGetStudentsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetStudentsStudentsPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("students", "get students", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetStudentsResponse returns a decoder for responses returned by the
// students get students endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeGetStudentsResponse may return the following errors:
//	- "internal_error" (type *students.CustomError): http.StatusInternalServerError
//	- error: internal error
func DecodeGetStudentsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetStudentsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("students", "get students", err)
			}
			p := NewGetStudentsStudentsOK(&body)
			view := "default"
			vres := &studentsviews.Students{Projected: p, View: view}
			if err = studentsviews.ValidateStudents(vres); err != nil {
				return nil, goahttp.ErrValidationError("students", "get students", err)
			}
			res := students.NewStudents(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body GetStudentsInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("students", "get students", err)
			}
			err = ValidateGetStudentsInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("students", "get students", err)
			}
			return nil, NewGetStudentsInternalError(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("students", "get students", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateStudentRequest instantiates a HTTP request object with method and
// path set to call the "students" service "create student" endpoint
func (c *Client) BuildCreateStudentRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateStudentStudentsPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("students", "create student", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeCreateStudentResponse returns a decoder for responses returned by the
// students create student endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeCreateStudentResponse may return the following errors:
//	- "internal_error" (type *students.CustomError): http.StatusInternalServerError
//	- "bad_request" (type *students.CustomError): http.StatusBadRequest
//	- error: internal error
func DecodeCreateStudentResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateStudentResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("students", "create student", err)
			}
			p := NewCreateStudentStudentOK(&body)
			view := "default"
			vres := &studentsviews.Student{Projected: p, View: view}
			if err = studentsviews.ValidateStudent(vres); err != nil {
				return nil, goahttp.ErrValidationError("students", "create student", err)
			}
			res := students.NewStudent(vres)
			return res, nil
		case http.StatusInternalServerError:
			var (
				body CreateStudentInternalErrorResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("students", "create student", err)
			}
			err = ValidateCreateStudentInternalErrorResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("students", "create student", err)
			}
			return nil, NewCreateStudentInternalError(&body)
		case http.StatusBadRequest:
			var (
				body CreateStudentBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("students", "create student", err)
			}
			err = ValidateCreateStudentBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("students", "create student", err)
			}
			return nil, NewCreateStudentBadRequest(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("students", "create student", resp.StatusCode, string(body))
		}
	}
}

// unmarshalStudentResponseBodyToStudentsviewsStudentView builds a value of
// type *studentsviews.StudentView from a value of type *StudentResponseBody.
func unmarshalStudentResponseBodyToStudentsviewsStudentView(v *StudentResponseBody) *studentsviews.StudentView {
	res := &studentsviews.StudentView{
		ID:             v.ID,
		Name:           v.Name,
		Ruby:           v.Ruby,
		StudentNumber:  v.StudentNumber,
		DateOfBirth:    v.DateOfBirth,
		Address:        v.Address,
		ExpirationDate: v.ExpirationDate,
	}

	return res
}
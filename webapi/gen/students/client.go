// Code generated by goa v3.7.6, DO NOT EDIT.
//
// students client
//
// Command:
// $ goa gen github.com/is-hoku/goa-template/webapi/design

package students

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "students" service client.
type Client struct {
	GetStudentEndpoint    goa.Endpoint
	GetStudentsEndpoint   goa.Endpoint
	CreateStudentEndpoint goa.Endpoint
}

// NewClient initializes a "students" service client given the endpoints.
func NewClient(getStudent, getStudents, createStudent goa.Endpoint) *Client {
	return &Client{
		GetStudentEndpoint:    getStudent,
		GetStudentsEndpoint:   getStudents,
		CreateStudentEndpoint: createStudent,
	}
}

// GetStudent calls the "get_student" endpoint of the "students" service.
// GetStudent may return the following errors:
//	- "internal_error" (type *CustomError)
//	- "not_found" (type *CustomError)
//	- error: internal error
func (c *Client) GetStudent(ctx context.Context, p *GetStudentPayload) (res *Student, err error) {
	var ires interface{}
	ires, err = c.GetStudentEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Student), nil
}

// GetStudents calls the "get_students" endpoint of the "students" service.
// GetStudents may return the following errors:
//	- "internal_error" (type *CustomError)
//	- error: internal error
func (c *Client) GetStudents(ctx context.Context) (res *Students, err error) {
	var ires interface{}
	ires, err = c.GetStudentsEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*Students), nil
}

// CreateStudent calls the "create_student" endpoint of the "students" service.
// CreateStudent may return the following errors:
//	- "internal_error" (type *CustomError)
//	- "bad_request" (type *CustomError)
//	- error: internal error
func (c *Client) CreateStudent(ctx context.Context) (res *Student, err error) {
	var ires interface{}
	ires, err = c.CreateStudentEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*Student), nil
}

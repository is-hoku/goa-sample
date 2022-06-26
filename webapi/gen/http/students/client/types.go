// Code generated by goa v3.7.6, DO NOT EDIT.
//
// students HTTP client types
//
// Command:
// $ goa gen github.com/is-hoku/goa-template/webapi/design

package client

import (
	students "github.com/is-hoku/goa-template/gen/students"
	studentsviews "github.com/is-hoku/goa-template/gen/students/views"
	goa "goa.design/goa/v3/pkg"
)

// GetStudentResponseBody is the type of the "students" service "get student"
// endpoint HTTP response body.
type GetStudentResponseBody struct {
	// 学生を一意に表す ID
	ID *int64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 学生の氏名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 学生の氏名のフリガナ
	Ruby *string `form:"ruby,omitempty" json:"ruby,omitempty" xml:"ruby,omitempty"`
	// 学生の学籍番号
	StudentNumber *int `form:"student_number,omitempty" json:"student_number,omitempty" xml:"student_number,omitempty"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth *string `form:"date_of_birth,omitempty" json:"date_of_birth,omitempty" xml:"date_of_birth,omitempty"`
	// 学生の住所
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate *string `form:"expiration_date,omitempty" json:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
}

// GetStudentsResponseBody is the type of the "students" service "get students"
// endpoint HTTP response body.
type GetStudentsResponseBody struct {
	Students []*StudentResponseBody `form:"students,omitempty" json:"students,omitempty" xml:"students,omitempty"`
}

// CreateStudentResponseBody is the type of the "students" service "create
// student" endpoint HTTP response body.
type CreateStudentResponseBody struct {
	// 学生を一意に表す ID
	ID *int64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 学生の氏名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 学生の氏名のフリガナ
	Ruby *string `form:"ruby,omitempty" json:"ruby,omitempty" xml:"ruby,omitempty"`
	// 学生の学籍番号
	StudentNumber *int `form:"student_number,omitempty" json:"student_number,omitempty" xml:"student_number,omitempty"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth *string `form:"date_of_birth,omitempty" json:"date_of_birth,omitempty" xml:"date_of_birth,omitempty"`
	// 学生の住所
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate *string `form:"expiration_date,omitempty" json:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
}

// GetStudentInternalErrorResponseBody is the type of the "students" service
// "get student" endpoint HTTP response body for the "internal_error" error.
type GetStudentInternalErrorResponseBody struct {
	// Name of error
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// GetStudentNotFoundResponseBody is the type of the "students" service "get
// student" endpoint HTTP response body for the "not_found" error.
type GetStudentNotFoundResponseBody struct {
	// Name of error
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// GetStudentsInternalErrorResponseBody is the type of the "students" service
// "get students" endpoint HTTP response body for the "internal_error" error.
type GetStudentsInternalErrorResponseBody struct {
	// Name of error
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateStudentInternalErrorResponseBody is the type of the "students" service
// "create student" endpoint HTTP response body for the "internal_error" error.
type CreateStudentInternalErrorResponseBody struct {
	// Name of error
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// CreateStudentBadRequestResponseBody is the type of the "students" service
// "create student" endpoint HTTP response body for the "bad_request" error.
type CreateStudentBadRequestResponseBody struct {
	// Name of error
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Message of error
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// StudentResponseBody is used to define fields on response body types.
type StudentResponseBody struct {
	// 学生を一意に表す ID
	ID *int64 `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// 学生の氏名
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// 学生の氏名のフリガナ
	Ruby *string `form:"ruby,omitempty" json:"ruby,omitempty" xml:"ruby,omitempty"`
	// 学生の学籍番号
	StudentNumber *int `form:"student_number,omitempty" json:"student_number,omitempty" xml:"student_number,omitempty"`
	// 学生の生年月日 (RFC3339)
	DateOfBirth *string `form:"date_of_birth,omitempty" json:"date_of_birth,omitempty" xml:"date_of_birth,omitempty"`
	// 学生の住所
	Address *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	// 学生証の有効期間 (RFC3339)
	ExpirationDate *string `form:"expiration_date,omitempty" json:"expiration_date,omitempty" xml:"expiration_date,omitempty"`
}

// NewGetStudentStudentOK builds a "students" service "get student" endpoint
// result from a HTTP "OK" response.
func NewGetStudentStudentOK(body *GetStudentResponseBody) *studentsviews.StudentView {
	v := &studentsviews.StudentView{
		ID:             body.ID,
		Name:           body.Name,
		Ruby:           body.Ruby,
		StudentNumber:  body.StudentNumber,
		DateOfBirth:    body.DateOfBirth,
		Address:        body.Address,
		ExpirationDate: body.ExpirationDate,
	}

	return v
}

// NewGetStudentInternalError builds a students service get student endpoint
// internal_error error.
func NewGetStudentInternalError(body *GetStudentInternalErrorResponseBody) *students.CustomError {
	v := &students.CustomError{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// NewGetStudentNotFound builds a students service get student endpoint
// not_found error.
func NewGetStudentNotFound(body *GetStudentNotFoundResponseBody) *students.CustomError {
	v := &students.CustomError{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// NewGetStudentsStudentsOK builds a "students" service "get students" endpoint
// result from a HTTP "OK" response.
func NewGetStudentsStudentsOK(body *GetStudentsResponseBody) *studentsviews.StudentsView {
	v := &studentsviews.StudentsView{}
	v.Students = make([]*studentsviews.StudentView, len(body.Students))
	for i, val := range body.Students {
		v.Students[i] = unmarshalStudentResponseBodyToStudentsviewsStudentView(val)
	}

	return v
}

// NewGetStudentsInternalError builds a students service get students endpoint
// internal_error error.
func NewGetStudentsInternalError(body *GetStudentsInternalErrorResponseBody) *students.CustomError {
	v := &students.CustomError{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// NewCreateStudentStudentOK builds a "students" service "create student"
// endpoint result from a HTTP "OK" response.
func NewCreateStudentStudentOK(body *CreateStudentResponseBody) *studentsviews.StudentView {
	v := &studentsviews.StudentView{
		ID:             body.ID,
		Name:           body.Name,
		Ruby:           body.Ruby,
		StudentNumber:  body.StudentNumber,
		DateOfBirth:    body.DateOfBirth,
		Address:        body.Address,
		ExpirationDate: body.ExpirationDate,
	}

	return v
}

// NewCreateStudentInternalError builds a students service create student
// endpoint internal_error error.
func NewCreateStudentInternalError(body *CreateStudentInternalErrorResponseBody) *students.CustomError {
	v := &students.CustomError{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// NewCreateStudentBadRequest builds a students service create student endpoint
// bad_request error.
func NewCreateStudentBadRequest(body *CreateStudentBadRequestResponseBody) *students.CustomError {
	v := &students.CustomError{
		Name:    *body.Name,
		Message: *body.Message,
	}

	return v
}

// ValidateGetStudentInternalErrorResponseBody runs the validations defined on
// get student_internal_error_response_body
func ValidateGetStudentInternalErrorResponseBody(body *GetStudentInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateGetStudentNotFoundResponseBody runs the validations defined on get
// student_not_found_response_body
func ValidateGetStudentNotFoundResponseBody(body *GetStudentNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateGetStudentsInternalErrorResponseBody runs the validations defined on
// get students_internal_error_response_body
func ValidateGetStudentsInternalErrorResponseBody(body *GetStudentsInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateCreateStudentInternalErrorResponseBody runs the validations defined
// on create student_internal_error_response_body
func ValidateCreateStudentInternalErrorResponseBody(body *CreateStudentInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateCreateStudentBadRequestResponseBody runs the validations defined on
// create student_bad_request_response_body
func ValidateCreateStudentBadRequestResponseBody(body *CreateStudentBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateStudentResponseBody runs the validations defined on
// StudentResponseBody
func ValidateStudentResponseBody(body *StudentResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Ruby == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ruby", "body"))
	}
	if body.StudentNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("student_number", "body"))
	}
	if body.DateOfBirth == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("date_of_birth", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.ExpirationDate == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("expiration_date", "body"))
	}
	if body.DateOfBirth != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.date_of_birth", *body.DateOfBirth, goa.FormatDateTime))
	}
	if body.ExpirationDate != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.expiration_date", *body.ExpirationDate, goa.FormatDateTime))
	}
	return
}
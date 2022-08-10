// Code generated by goa v3.7.6, DO NOT EDIT.
//
// HTTP request path constructors for the student service.
//
// Command:
// $ goa gen github.com/is-hoku/goa-sample/webapi/design

package server

import (
	"fmt"
)

// GetStudentStudentPath returns the URL path to the student service get_student HTTP endpoint.
func GetStudentStudentPath(studentNumber uint32) string {
	return fmt.Sprintf("/students/%v", studentNumber)
}

// GetStudentsStudentPath returns the URL path to the student service get_students HTTP endpoint.
func GetStudentsStudentPath() string {
	return "/students"
}

// CreateStudentStudentPath returns the URL path to the student service create_student HTTP endpoint.
func CreateStudentStudentPath() string {
	return "/students"
}

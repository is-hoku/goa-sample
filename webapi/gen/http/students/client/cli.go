// Code generated by goa v3.7.6, DO NOT EDIT.
//
// students HTTP client CLI support package
//
// Command:
// $ goa gen github.com/is-hoku/goa-template/webapi/design

package client

import (
	"fmt"
	"strconv"

	students "github.com/is-hoku/goa-template/gen/students"
)

// BuildGetStudentPayload builds the payload for the students get student
// endpoint from CLI flags.
func BuildGetStudentPayload(studentsGetStudentID string) (*students.GetStudentPayload, error) {
	var err error
	var id int64
	{
		id, err = strconv.ParseInt(studentsGetStudentID, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT64")
		}
	}
	v := &students.GetStudentPayload{}
	v.ID = &id

	return v, nil
}
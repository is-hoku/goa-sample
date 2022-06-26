// Code generated by goa v3.7.6, DO NOT EDIT.
//
// students service
//
// Command:
// $ goa gen github.com/is-hoku/goa-template/webapi/design

package students

import (
	"context"

	studentsviews "github.com/is-hoku/goa-template/gen/students/views"
)

// Service is the students service interface.
type Service interface {
	// id から学生を取得する。
	GetStudent(context.Context, *GetStudentPayload) (res *Student, err error)
	// 学籍番号で昇順にソートされた全ての学生を取得する。
	GetStudents(context.Context) (res *Students, err error)
	// 学生を登録する。
	CreateStudent(context.Context) (res *Student, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "students"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [3]string{"get student", "get students", "create student"}

type CustomError struct {
	// Name of error
	Name string
	// Message of error
	Message string
}

// GetStudentPayload is the payload type of the students service get student
// method.
type GetStudentPayload struct {
	// Student's unique ID
	ID *int64
}

// Student is the result type of the students service get student method.
type Student struct {
	// 学生を一意に表す ID
	ID int64
	// 学生の氏名
	Name string
	// 学生の氏名のフリガナ
	Ruby string
	// 学生の学籍番号
	StudentNumber int
	// 学生の生年月日 (RFC3339)
	DateOfBirth string
	// 学生の住所
	Address string
	// 学生証の有効期間 (RFC3339)
	ExpirationDate string
}

// Students is the result type of the students service get students method.
type Students struct {
	Students []*Student
}

// Error returns an error description.
func (e *CustomError) Error() string {
	return ""
}

// ErrorName returns "CustomError".
func (e *CustomError) ErrorName() string {
	return e.Name
}

// NewStudent initializes result type Student from viewed result type Student.
func NewStudent(vres *studentsviews.Student) *Student {
	return newStudent(vres.Projected)
}

// NewViewedStudent initializes viewed result type Student from result type
// Student using the given view.
func NewViewedStudent(res *Student, view string) *studentsviews.Student {
	p := newStudentView(res)
	return &studentsviews.Student{Projected: p, View: "default"}
}

// NewStudents initializes result type Students from viewed result type
// Students.
func NewStudents(vres *studentsviews.Students) *Students {
	return newStudents(vres.Projected)
}

// NewViewedStudents initializes viewed result type Students from result type
// Students using the given view.
func NewViewedStudents(res *Students, view string) *studentsviews.Students {
	p := newStudentsView(res)
	return &studentsviews.Students{Projected: p, View: "default"}
}

// newStudent converts projected type Student to service type Student.
func newStudent(vres *studentsviews.StudentView) *Student {
	res := &Student{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Ruby != nil {
		res.Ruby = *vres.Ruby
	}
	if vres.StudentNumber != nil {
		res.StudentNumber = *vres.StudentNumber
	}
	if vres.DateOfBirth != nil {
		res.DateOfBirth = *vres.DateOfBirth
	}
	if vres.Address != nil {
		res.Address = *vres.Address
	}
	if vres.ExpirationDate != nil {
		res.ExpirationDate = *vres.ExpirationDate
	}
	return res
}

// newStudentView projects result type Student to projected type StudentView
// using the "default" view.
func newStudentView(res *Student) *studentsviews.StudentView {
	vres := &studentsviews.StudentView{
		ID:             &res.ID,
		Name:           &res.Name,
		Ruby:           &res.Ruby,
		StudentNumber:  &res.StudentNumber,
		DateOfBirth:    &res.DateOfBirth,
		Address:        &res.Address,
		ExpirationDate: &res.ExpirationDate,
	}
	return vres
}

// newStudents converts projected type Students to service type Students.
func newStudents(vres *studentsviews.StudentsView) *Students {
	res := &Students{}
	if vres.Students != nil {
		res.Students = make([]*Student, len(vres.Students))
		for i, val := range vres.Students {
			res.Students[i] = transformStudentsviewsStudentViewToStudent(val)
		}
	}
	return res
}

// newStudentsView projects result type Students to projected type StudentsView
// using the "default" view.
func newStudentsView(res *Students) *studentsviews.StudentsView {
	vres := &studentsviews.StudentsView{}
	if res.Students != nil {
		vres.Students = make([]*studentsviews.StudentView, len(res.Students))
		for i, val := range res.Students {
			vres.Students[i] = transformStudentToStudentsviewsStudentView(val)
		}
	}
	return vres
}

// transformStudentsviewsStudentViewToStudent builds a value of type *Student
// from a value of type *studentsviews.StudentView.
func transformStudentsviewsStudentViewToStudent(v *studentsviews.StudentView) *Student {
	if v == nil {
		return nil
	}
	res := &Student{
		ID:             *v.ID,
		Name:           *v.Name,
		Ruby:           *v.Ruby,
		StudentNumber:  *v.StudentNumber,
		DateOfBirth:    *v.DateOfBirth,
		Address:        *v.Address,
		ExpirationDate: *v.ExpirationDate,
	}

	return res
}

// transformStudentToStudentsviewsStudentView builds a value of type
// *studentsviews.StudentView from a value of type *Student.
func transformStudentToStudentsviewsStudentView(v *Student) *studentsviews.StudentView {
	res := &studentsviews.StudentView{
		ID:             &v.ID,
		Name:           &v.Name,
		Ruby:           &v.Ruby,
		StudentNumber:  &v.StudentNumber,
		DateOfBirth:    &v.DateOfBirth,
		Address:        &v.Address,
		ExpirationDate: &v.ExpirationDate,
	}

	return res
}
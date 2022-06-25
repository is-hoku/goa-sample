package studentsapi

import (
	"context"
	"log"

	students "github.com/is-hoku/goa-template/gen/students"
)

// students service example implementation.
// The example methods log the requests and return zero values.
type studentssrvc struct {
	logger *log.Logger
}

// NewStudents returns the students service implementation.
func NewStudents(logger *log.Logger) students.Service {
	return &studentssrvc{logger}
}

// id から学生を取得する。
func (s *studentssrvc) GetStudent(ctx context.Context, p *students.GetStudentPayload) (res *students.Student, err error) {
	res = &students.Student{}
	s.logger.Print("students.get student")
	return
}

// 学籍番号で昇順にソートされた全ての学生を取得する。
func (s *studentssrvc) GetStudents(ctx context.Context) (res *students.Students, err error) {
	res = &students.Students{}
	s.logger.Print("students.get students")
	return
}

// 学生を登録する。
func (s *studentssrvc) CreateStudent(ctx context.Context) (res *students.Student, err error) {
	res = &students.Student{}
	s.logger.Print("students.create student")
	return
}

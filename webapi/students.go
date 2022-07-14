package studentsapi

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/is-hoku/goa-template/webapi/datastore"
	students "github.com/is-hoku/goa-template/webapi/gen/students"
	"github.com/is-hoku/goa-template/webapi/interactor"
	"github.com/is-hoku/goa-template/webapi/model"
	"github.com/is-hoku/goa-template/webapi/repository"
	"github.com/joho/godotenv"
)

// students service example implementation.
// The example methods log the requests and return zero values.
type studentssrvc struct {
	logger  *log.Logger
	handler repository.StudentRepository
}

// NewStudents returns the students service implementation.
func NewStudents(logger *log.Logger) students.Service {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Cannot load .env: %v", err)
	}
	config := &datastore.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
	handler, err := datastore.New(config)
	return &studentssrvc{logger, handler}
}

// id から学生を取得する。
func (s *studentssrvc) GetStudent(ctx context.Context, p *students.GetStudentPayload) (res *students.Student, err error) {
	s.logger.Print("students.get student")
	si := interactor.StudentInteractor{Repo: s.handler}
	student, err := si.GetByNum(ctx, *p.StudentNumber)
	if err != nil {
		return nil, err
	}
	res = &students.Student{
		ID:             student.ID,
		Name:           student.Name,
		Ruby:           student.Ruby,
		StudentNumber:  student.StudentNumber,
		DateOfBirth:    student.DateOfBirth.String(),
		Address:        student.Address,
		ExpirationDate: student.ExpirationDate.String(),
	}
	return
}

// 学籍番号で昇順にソートされた全ての学生を取得する。
func (s *studentssrvc) GetStudents(ctx context.Context) (res *students.Students, err error) {
	s.logger.Print("students.get students")
	si := interactor.StudentInteractor{Repo: s.handler}
	allStudents, err := si.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	l := make([]*students.Student, 0, len(allStudents))
	for _, student := range allStudents {
		var st *students.Student
		st = &students.Student{
			ID:             student.ID,
			Name:           student.Name,
			Ruby:           student.Ruby,
			StudentNumber:  student.StudentNumber,
			DateOfBirth:    student.DateOfBirth.String(),
			Address:        student.Address,
			ExpirationDate: student.ExpirationDate.String(),
		}
		l = append(l, st)
	}
	res = &students.Students{Students: l}
	return
}

// 学生を登録する。
func (s *studentssrvc) CreateStudent(ctx context.Context, body *students.StudentBody) (res *students.Student, err error) {
	s.logger.Print("students.create student")
	si := interactor.StudentInteractor{Repo: s.handler}
	birth, err := time.Parse(time.RFC3339, body.DateOfBirth)
	if err != nil {
		return nil, err
	}
	expiration, err := time.Parse(time.RFC3339, body.ExpirationDate)
	if err != nil {
		return nil, err
	}
	bodyStudent := &model.Student{
		Name:           body.Name,
		Ruby:           body.Ruby,
		StudentNumber:  body.StudentNumber,
		DateOfBirth:    birth,
		Address:        body.Address,
		ExpirationDate: expiration,
	}
	student, err := si.Create(ctx, bodyStudent)
	if err != nil {
		return nil, err
	}
	res = &students.Student{
		ID:             student.ID,
		Name:           student.Name,
		Ruby:           student.Ruby,
		StudentNumber:  student.StudentNumber,
		DateOfBirth:    student.DateOfBirth.String(),
		Address:        student.Address,
		ExpirationDate: student.ExpirationDate.String(),
	}
	return
}

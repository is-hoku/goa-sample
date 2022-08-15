package studentsapi

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/is-hoku/goa-sample/webapi/datastore"
	"github.com/is-hoku/goa-sample/webapi/gen/student"
	"github.com/is-hoku/goa-sample/webapi/interactor"
	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/repository"
	"github.com/joho/godotenv"
)

// students service example implementation.
// The example methods log the requests and return zero values.
type studentsrvc struct {
	logger  *log.Logger
	handler repository.StudentRepository
}

// NewStudents returns the students service implementation.
func NewStudent(logger *log.Logger) student.Service {
	err := godotenv.Load("../.env")
	if err != nil {
		logger.Fatalf("Could not load .env: %s", err)
	}
	config := &datastore.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
	sqldb, err := datastore.NewDB(config)
	if err != nil {
		logger.Fatalf("Could not generate db: %s", err)
	}
	studentDB := &datastore.StudentHandler{DBHandler: sqldb}
	return &studentsrvc{logger, studentDB}
}

// 学籍番号から学生を取得する。
func (s *studentsrvc) GetStudent(ctx context.Context, p *student.GetStudentPayload) (*student.Student, error) {
	s.logger.Print("students.get student")
	si := interactor.StudentInteractor{Repo: s.handler}
	gotStudent, err := si.GetByNumber(ctx, *p.StudentNumber)
	if err != nil {
		return nil, err
	}
	res := &student.Student{
		ID:             gotStudent.ID,
		Name:           gotStudent.Name,
		Ruby:           gotStudent.Ruby,
		StudentNumber:  gotStudent.StudentNumber,
		DateOfBirth:    gotStudent.DateOfBirth.Format(time.RFC3339),
		Address:        gotStudent.Address,
		ExpirationDate: gotStudent.ExpirationDate.Format(time.RFC3339),
	}
	return res, nil
}

// 学籍番号で昇順にソートされた全ての学生を取得する。
func (s *studentsrvc) GetStudents(ctx context.Context) (*student.Students, error) {
	s.logger.Print("students.get students")
	si := interactor.StudentInteractor{Repo: s.handler}
	allStudents, err := si.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	l := make([]*student.Student, 0, len(allStudents))
	for _, person := range allStudents {
		var st *student.Student
		st = &student.Student{
			ID:             person.ID,
			Name:           person.Name,
			Ruby:           person.Ruby,
			StudentNumber:  person.StudentNumber,
			DateOfBirth:    person.DateOfBirth.Format(time.RFC3339),
			Address:        person.Address,
			ExpirationDate: person.ExpirationDate.Format(time.RFC3339),
		}
		l = append(l, st)
	}
	res := &student.Students{Students: l}
	return res, nil
}

// 学生を登録する。
func (s *studentsrvc) CreateStudent(ctx context.Context, body *student.StudentBody) (*student.Student, error) {
	s.logger.Print("students.create student")
	si := interactor.StudentInteractor{Repo: s.handler}
	birth, err := time.Parse(time.RFC3339, body.DateOfBirth)
	if err != nil {
		return nil, &student.CustomError{Name: "bad_request", Message: "date_of_birth is invalid format"}
	}
	expiration, err := time.Parse(time.RFC3339, body.ExpirationDate)
	if err != nil {
		return nil, &student.CustomError{Name: "bad_request", Message: "expiration_date is invalid format"}
	}
	bodyStudent := &model.Student{
		Name:           body.Name,
		Ruby:           body.Ruby,
		StudentNumber:  body.StudentNumber,
		DateOfBirth:    birth,
		Address:        body.Address,
		ExpirationDate: expiration,
	}
	createdStudent, err := si.Create(ctx, bodyStudent)
	if err != nil {
		return nil, err
	}
	res := &student.Student{
		ID:             createdStudent.ID,
		Name:           createdStudent.Name,
		Ruby:           createdStudent.Ruby,
		StudentNumber:  createdStudent.StudentNumber,
		DateOfBirth:    createdStudent.DateOfBirth.Format(time.RFC3339),
		Address:        createdStudent.Address,
		ExpirationDate: createdStudent.ExpirationDate.Format(time.RFC3339),
	}
	return res, nil
}

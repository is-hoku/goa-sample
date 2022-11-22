package sample

import (
	"context"
	"log"
	"time"

	"firebase.google.com/go/auth"
	student "github.com/is-hoku/goa-sample/webapi/gen/student"
	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/usecase"
	"goa.design/goa/v3/security"
)

type studentsrvc struct {
	logger *log.Logger
	app    *StudentApp
}

func NewStudent(logger *log.Logger, app *StudentApp) student.Service {
	return &studentsrvc{logger, app}
}

func (s *studentsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	// ダミーユーザ
	switch token {
	case "test":
		testUser := &auth.UserInfo{
			DisplayName: "test",
			Email:       "test@example.com",
			PhoneNumber: "000-0000-0000",
			PhotoURL:    "photo.example.com",
			ProviderID:  "testprovider",
			UID:         "abcdefg",
		}
		ctx = SetTestUserInfo(ctx, testUser)
		return ctx, nil
	default:
		// Firebase Authentication で認証
		authClient, err := newFirebase(ctx)
		if err != nil {
			return ctx, &student.CustomError{Name: "internal_error", Message: "Failed to create firebase app instance"}
		}
		idToken, err := authClient.VerifyIDToken(ctx, token)
		if err != nil {
			return ctx, &student.CustomError{Name: "unauthorized", Message: "Unauthorized"}
		}
		ctx, err = SetUserInfo(ctx, idToken)
		if err != nil {
			return ctx, &student.CustomError{Name: "internal_error", Message: "Failed to get user information from Firebase"}
		}
		return ctx, nil
	}
}

// 学籍番号から学生を取得する。
func (s *studentsrvc) GetStudent(ctx context.Context, p *student.GetStudentPayload) (*student.Student, error) {
	s.logger.Print("students.get student")
	got, err := s.app.getter.GetStudentByNumber(ctx, &usecase.GetStudentByNumberInput{
		StudentNumber: *p.StudentNumber,
	})
	if err != nil {
		return nil, err
	}
	res := &student.Student{
		ID:             got.Student.ID,
		Name:           got.Student.Name,
		Ruby:           got.Student.Ruby,
		StudentNumber:  got.Student.StudentNumber,
		DateOfBirth:    got.Student.DateOfBirth.Format(time.RFC3339),
		Address:        got.Student.Address,
		ExpirationDate: got.Student.ExpirationDate.Format(time.RFC3339),
	}
	return res, nil
}

// 学籍番号で昇順にソートされた全ての学生を取得する。
func (s *studentsrvc) GetStudents(ctx context.Context, p *student.GetStudentsPayload) (*student.Students, error) {
	s.logger.Print("students.get students")
	students, err := s.app.multigetter.GetStudents(ctx)
	if err != nil {
		return nil, err
	}
	l := make([]*student.Student, 0, len(students.Students))
	for _, got := range students.Students {
		st := &student.Student{
			ID:             got.ID,
			Name:           got.Name,
			Ruby:           got.Ruby,
			StudentNumber:  got.StudentNumber,
			DateOfBirth:    got.DateOfBirth.Format(time.RFC3339),
			Address:        got.Address,
			ExpirationDate: got.ExpirationDate.Format(time.RFC3339),
		}
		l = append(l, st)
	}
	res := &student.Students{Students: l}
	return res, nil
}

// 学生を登録する。
func (s *studentsrvc) CreateStudent(ctx context.Context, body *student.CreateStudentPayload) (*student.Student, error) {
	s.logger.Print("students.create student")
	birth, err := time.Parse(time.RFC3339, body.DateOfBirth)
	if err != nil {
		return nil, &student.CustomError{Name: "bad_request", Message: "date_of_birth is invalid format"}
	}
	expiration, err := time.Parse(time.RFC3339, body.ExpirationDate)
	if err != nil {
		return nil, &student.CustomError{Name: "bad_request", Message: "expiration_date is invalid format"}
	}
	bodyStudent := &usecase.CreateStudentInput{
		Student: &model.Student{
			Name:           body.Name,
			Ruby:           body.Ruby,
			StudentNumber:  body.StudentNumber,
			DateOfBirth:    birth,
			Address:        body.Address,
			ExpirationDate: expiration,
		},
	}
	created, err := s.app.creator.CreateStudent(ctx, bodyStudent)
	if err != nil {
		return nil, err
	}
	res := &student.Student{
		ID:             created.Student.ID,
		Name:           created.Student.Name,
		Ruby:           created.Student.Ruby,
		StudentNumber:  created.Student.StudentNumber,
		DateOfBirth:    created.Student.DateOfBirth.Format(time.RFC3339),
		Address:        created.Student.Address,
		ExpirationDate: created.Student.ExpirationDate.Format(time.RFC3339),
	}
	return res, nil
}

// 学生情報を更新する。
func (s *studentsrvc) UpdateStudent(ctx context.Context, body *student.UpdateStudentPayload) (*student.Student, error) {
	s.logger.Print("students.update student")
	return nil, nil
}

// 学生を削除する。
func (s *studentsrvc) DeleteStudent(ctx context.Context, p *student.DeleteStudentPayload) error {
	s.logger.Print("students.delete student")
	return nil
}

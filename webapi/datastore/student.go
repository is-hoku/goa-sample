package datastore

import (
	"context"

	"github.com/is-hoku/goa-sample/webapi/gen/student"
	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/repository"
)

var _ repository.StudentRepository = (*StudentHandler)(nil)

type StudentHandler struct {
	*DBHandler
}

func (db *StudentHandler) GetByNumber(ctx context.Context, number uint32) (*model.Student, error) {
	queries := New(db.DB)
	s, err := queries.GetStudentByNumber(ctx, number)
	if err != nil {
		if err != nil {
			return nil, &student.CustomError{Name: "not_found", Message: "Student Not Found"}
		}
		return nil, err
	}
	res := &model.Student{
		ID:             s.ID,
		Name:           s.Name,
		Ruby:           s.Ruby,
		StudentNumber:  s.StudentNumber,
		DateOfBirth:    s.DateOfBirth,
		Address:        s.Address,
		ExpirationDate: s.ExpirationDate,
	}
	return res, nil
}

func (db *StudentHandler) GetByID(ctx context.Context, id uint64) (*model.Student, error) {
	queries := New(db.DB)
	s, err := queries.GetStudentByID(ctx, id)
	if err != nil {
		return nil, &student.CustomError{Name: "not_found", Message: "Student Not Found"}
	}
	res := &model.Student{
		ID:             s.ID,
		Name:           s.Name,
		Ruby:           s.Ruby,
		StudentNumber:  s.StudentNumber,
		DateOfBirth:    s.DateOfBirth,
		Address:        s.Address,
		ExpirationDate: s.ExpirationDate,
	}
	return res, nil
}

func (db *StudentHandler) Set(ctx context.Context, s *model.Student) (uint64, error) {
	queries := New(db.DB)
	params := SetStudentParams{
		Name:           s.Name,
		Ruby:           s.Ruby,
		StudentNumber:  s.StudentNumber,
		DateOfBirth:    s.DateOfBirth,
		Address:        s.Address,
		ExpirationDate: s.ExpirationDate,
	}
	result, err := queries.SetStudent(ctx, params)
	if err != nil {
		return 0, &student.CustomError{Name: "bad_request", Message: "Bad Request Body"}
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		return 0, &student.CustomError{Name: "internal_error", Message: "Internal Server Error"}
	}
	return uint64(insertedID), nil
}

func (db *StudentHandler) GetAll(ctx context.Context) ([]*model.Student, error) {
	queries := New(db.DB)
	students, err := queries.GetAllStudents(ctx)
	if err != nil {
		return nil, &student.CustomError{Name: "internal_error", Message: "Internal Server Error"}
	}
	var res []*model.Student
	for _, s := range students {
		student := &model.Student{
			ID:             s.ID,
			Name:           s.Name,
			Ruby:           s.Ruby,
			StudentNumber:  s.StudentNumber,
			DateOfBirth:    s.DateOfBirth,
			Address:        s.Address,
			ExpirationDate: s.ExpirationDate,
		}
		res = append(res, student)
	}
	return res, nil
}

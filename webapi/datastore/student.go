package datastore

import (
	"context"
	"errors"

	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/repository"
)

var _ repository.StudentRepository = (*StudentHandler)(nil)

type StudentHandler struct {
	*DBHandler
}

func (db *StudentHandler) GetByNumber(ctx context.Context, number int32) (*model.Student, error) {
	queries := New(db.DB)
	s, err := queries.GetStudentByNumber(ctx, number)
	if err != nil {
		if err != nil {
			return nil, errors.New("not_found")
		}
		return nil, err
	}
	res := &model.Student{
		ID:             uint64(s.ID),
		Name:           s.Name,
		Ruby:           s.Ruby,
		StudentNumber:  uint32(s.StudentNumber),
		DateOfBirth:    s.DateOfBirth,
		Address:        s.Address,
		ExpirationDate: s.ExpirationDate,
	}
	return res, nil
}

func (db *StudentHandler) GetByID(ctx context.Context, id int64) (*model.Student, error) {
	queries := New(db.DB)
	s, err := queries.GetStudentByID(ctx, id)
	if err != nil {
		return nil, errors.New("not_found")
	}
	res := &model.Student{
		ID:             uint64(s.ID),
		Name:           s.Name,
		Ruby:           s.Ruby,
		StudentNumber:  uint32(s.StudentNumber),
		DateOfBirth:    s.DateOfBirth,
		Address:        s.Address,
		ExpirationDate: s.ExpirationDate,
	}
	return res, nil
}

func (db *StudentHandler) Set(ctx context.Context, student *model.Student) (int64, error) {
	queries := New(db.DB)
	params := SetStudentParams{
		Name:           student.Name,
		Ruby:           student.Ruby,
		StudentNumber:  int32(student.StudentNumber),
		DateOfBirth:    student.DateOfBirth,
		Address:        student.Address,
		ExpirationDate: student.ExpirationDate,
	}
	result, err := queries.SetStudent(ctx, params)
	if err != nil {
		return -1, errors.New("bad_request")
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		return -1, errors.New("internal_error")
	}
	return insertedID, nil
}

func (db *StudentHandler) GetAll(ctx context.Context) ([]*model.Student, error) {
	queries := New(db.DB)
	students, err := queries.GetAllStudents(ctx)
	if err != nil {
		return nil, errors.New("internal_error")
	}
	var res []*model.Student
	for _, s := range students {
		student := &model.Student{
			ID:             uint64(s.ID),
			Name:           s.Name,
			Ruby:           s.Ruby,
			StudentNumber:  uint32(s.StudentNumber),
			DateOfBirth:    s.DateOfBirth,
			Address:        s.Address,
			ExpirationDate: s.ExpirationDate,
		}
		res = append(res, student)
	}
	return res, nil
}

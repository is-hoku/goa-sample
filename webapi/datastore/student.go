package datastore

import (
	"context"

	"github.com/is-hoku/goa-sample/webapi/gen/student"
	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/repository"
)

var _ repository.StudentByNumberGetter = (*GetStudentByNumberMedia)(nil)

type GetStudentByNumberMedia struct {
	*DB
}

func NewGetStudentByNumberMedia(db *DB) *GetStudentByNumberMedia {
	return &GetStudentByNumberMedia{db}
}

func (db *GetStudentByNumberMedia) GetStudentByNumber(ctx context.Context, input *repository.GetStudentByNumberInput) (*repository.GetStudentByNumberOutput, error) {
	queries := New(db.DB)
	s, err := queries.GetStudentByNumber(ctx, input.StudentNumber)
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
	return &repository.GetStudentByNumberOutput{
		Student: res,
	}, nil
}

var _ repository.StudentByIDGetter = (*GetStudentByIDMedia)(nil)

type GetStudentByIDMedia struct {
	*DB
}

func NewGetStudentByIDMedia(db *DB) *GetStudentByIDMedia {
	return &GetStudentByIDMedia{db}
}

func (db *GetStudentByIDMedia) GetStudentByID(ctx context.Context, input *repository.GetStudentByIDInput) (*repository.GetStudentByIDOutput, error) {
	queries := New(db.DB)
	s, err := queries.GetStudentByID(ctx, input.ID)
	if err != nil {
		return nil, &student.CustomError{Name: "not_found", Message: "Student Not Found"}
	}
	return &repository.GetStudentByIDOutput{
		Student: &model.Student{
			ID:             s.ID,
			Name:           s.Name,
			Ruby:           s.Ruby,
			StudentNumber:  s.StudentNumber,
			DateOfBirth:    s.DateOfBirth,
			Address:        s.Address,
			ExpirationDate: s.ExpirationDate,
		},
	}, nil
}

var _ repository.StudentCreator = (*CreateStudentMedia)(nil)

type CreateStudentMedia struct {
	*DB
}

func NewCreateStudentMedia(db *DB) *CreateStudentMedia {
	return &CreateStudentMedia{db}
}

func (db *CreateStudentMedia) CreateStudent(ctx context.Context, input *repository.CreateStudentInput) (*repository.CreateStudentOutput, error) {
	queries := New(db.DB)
	params := SetStudentParams{
		Name:           input.Student.Name,
		Ruby:           input.Student.Ruby,
		StudentNumber:  input.Student.StudentNumber,
		DateOfBirth:    input.Student.DateOfBirth,
		Address:        input.Student.Address,
		ExpirationDate: input.Student.ExpirationDate,
	}
	result, err := queries.SetStudent(ctx, params)
	if err != nil {
		return nil, &student.CustomError{Name: "bad_request", Message: "Bad Request Body"}
	}
	insertedID, err := result.LastInsertId()
	if err != nil {
		return nil, &student.CustomError{Name: "internal_error", Message: "Internal Server Error"}
	}
	return &repository.CreateStudentOutput{
		ID: uint64(insertedID),
	}, nil
}

var _ repository.StudentsGetter = (*GetStudentsMedia)(nil)

type GetStudentsMedia struct {
	*DB
}

func NewGetStudentsMedia(db *DB) *GetStudentsMedia {
	return &GetStudentsMedia{db}
}
func (db *GetStudentsMedia) GetStudents(ctx context.Context) (*repository.GetStudentsOutput, error) {
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
	return &repository.GetStudentsOutput{
		Students: res,
	}, nil
}

var _ repository.StudentUpdater = (*UpdateStudentMedia)(nil)

type UpdateStudentMedia struct {
	*DB
}

func NewUpdateStudentMedia(db *DB) *UpdateStudentMedia {
	return &UpdateStudentMedia{db}
}

func (db *UpdateStudentMedia) UpdateStudent(ctx context.Context, input *repository.UpdateStudentInput) (*repository.UpdateStudentOutput, error) {
	queries := New(db.DB)
	params := UpdateStudentParams{
		ID:             input.Student.ID,
		Name:           input.Student.Name,
		Ruby:           input.Student.Ruby,
		StudentNumber:  input.Student.StudentNumber,
		DateOfBirth:    input.Student.DateOfBirth,
		Address:        input.Student.Address,
		ExpirationDate: input.Student.ExpirationDate,
	}
	result, err := queries.UpdateStudent(ctx, params)
	if err != nil {
		return nil, &student.CustomError{Name: "bad_request", Message: "Bad Request Body"}
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return nil, &student.CustomError{Name: "internal_error", Message: "Internal Server Error"}
	}
	if affected == 0 {
		return nil, &student.CustomError{Name: "not_found", Message: "Student Not Found"}
	}
	return &repository.UpdateStudentOutput{
		ID: input.Student.ID,
	}, nil
}

var _ repository.StudentDeleter = (*DeleteStudentMedia)(nil)

type DeleteStudentMedia struct {
	*DB
}

func NewDeleteStudentMedia(db *DB) *DeleteStudentMedia {
	return &DeleteStudentMedia{db}
}

func (db *DeleteStudentMedia) DeleteStudent(ctx context.Context, input *repository.DeleteStudentInput) error {
	queries := New(db.DB)
	result, err := queries.DeleteStudent(ctx, input.StudentNumber)
	if err != nil {
		return &student.CustomError{Name: "bad_request", Message: "Bad Request Body"}
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return &student.CustomError{Name: "internal_error", Message: "Internal Server Error"}
	}
	if affected == 0 {
		return &student.CustomError{Name: "not_found", Message: "Student Not Found"}
	}
	return nil
}

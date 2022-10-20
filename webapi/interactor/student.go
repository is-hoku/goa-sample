package interactor

import (
	"context"

	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/repository"
	"github.com/is-hoku/goa-sample/webapi/usecase"
)

var _ usecase.StudentByNumberGetter = (*GetStudentByNumber)(nil)

type GetStudentByNumber struct {
	opt *GetStudentByNumberOption
}

type GetStudentByNumberOption struct {
	repository.StudentByNumberGetter
}

func NewGetStudentByNumber(opt *GetStudentByNumberOption) *GetStudentByNumber {
	return &GetStudentByNumber{
		opt: opt,
	}
}

func (s *GetStudentByNumber) GetStudentByNumber(ctx context.Context, input *usecase.GetStudentByNumberInput) (*usecase.GetStudentByNumberOutput, error) {
	student, err := s.opt.GetStudentByNumber(ctx, &repository.GetStudentByNumberInput{
		StudentNumber: input.StudentNumber,
	})
	if err != nil {
		return nil, err
	}
	return &usecase.GetStudentByNumberOutput{
		Student: &model.Student{
			ID:             student.Student.ID,
			Name:           student.Student.Name,
			Ruby:           student.Student.Ruby,
			StudentNumber:  student.Student.StudentNumber,
			DateOfBirth:    student.Student.DateOfBirth,
			Address:        student.Student.Address,
			ExpirationDate: student.Student.ExpirationDate,
		},
	}, nil
}

var _ usecase.StudentCreator = (*CreateStudent)(nil)

type CreateStudent struct {
	opt *CreateStudentOption
}

type CreateStudentOption struct {
	repository.StudentByIDGetter
	repository.StudentCreator
}

func NewCreateStudent(opt *CreateStudentOption) *CreateStudent {
	return &CreateStudent{
		opt: opt,
	}
}

func (s *CreateStudent) CreateStudent(ctx context.Context, input *usecase.CreateStudentInput) (*usecase.CreateStudentOutput, error) {
	tx, err := s.opt.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	studentID, err := s.opt.CreateStudent(ctx, &repository.CreateStudentInput{
		Student: input.Student,
	})
	if err != nil {
		return nil, err
	}
	student, err := s.opt.GetStudentByID(ctx, &repository.GetStudentByIDInput{
		ID: studentID.ID,
	})
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &usecase.CreateStudentOutput{
		Student: student.Student,
	}, nil
}

var _ usecase.StudentsGetter = (*GetStudents)(nil)

type GetStudents struct {
	opt *GetStudentsOption
}

type GetStudentsOption struct {
	repository.StudentsGetter
}

func NewGetStudents(opt *GetStudentsOption) *GetStudents {
	return &GetStudents{
		opt: opt,
	}
}
func (s *GetStudents) GetStudents(ctx context.Context) (*usecase.GetStudentsOutput, error) {
	gotStudents, err := s.opt.GetStudents(ctx)
	if err != nil {
		return nil, err
	}
	var students []*model.Student
	for _, s := range gotStudents.Students {
		students = append(students, s)
	}
	return &usecase.GetStudentsOutput{
		Students: students,
	}, nil
}

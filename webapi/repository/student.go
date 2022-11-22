package repository

import (
	"context"

	"github.com/is-hoku/goa-sample/webapi/model"
)

type StudentByNumberGetter interface {
	GetStudentByNumber(ctx context.Context, input *GetStudentByNumberInput) (*GetStudentByNumberOutput, error)
}

type GetStudentByNumberInput struct {
	StudentNumber uint32
}

type GetStudentByNumberOutput struct {
	Student *model.Student
}

type StudentByIDGetter interface {
	GetStudentByID(ctx context.Context, input *GetStudentByIDInput) (*GetStudentByIDOutput, error)
}

type GetStudentByIDInput struct {
	ID uint64
}

type GetStudentByIDOutput struct {
	Student *model.Student
}

type StudentsGetter interface {
	GetStudents(ctx context.Context) (*GetStudentsOutput, error)
}

type GetStudentsOutput struct {
	Students []*model.Student
}

type StudentCreator interface {
	CreateStudent(ctx context.Context, input *CreateStudentInput) (*CreateStudentOutput, error)
	TxBeginner
}

type CreateStudentInput struct {
	Student *model.Student
}

type CreateStudentOutput struct {
	ID uint64
}

type StudentUpdater interface {
	UpdateStudent(ctx context.Context, input *UpdateStudentInput) (*UpdateStudentOutput, error)
	TxBeginner
}

type UpdateStudentInput struct {
	Student *model.Student
}

type UpdateStudentOutput struct {
	ID uint64
}

type StudentDeleter interface {
	DeleteStudent(ctx context.Context, input *DeleteStudentInput) error
}

type DeleteStudentInput struct {
	StudentNumber uint32
}

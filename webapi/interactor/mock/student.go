package mock

import (
	"context"

	"github.com/is-hoku/goa-sample/webapi/usecase"
)

var _ usecase.StudentByNumberGetter = GetStudentByNumber(nil)

type GetStudentByNumber func(ctx context.Context, input *usecase.GetStudentByNumberInput) (*usecase.GetStudentByNumberOutput, error)

func (mock GetStudentByNumber) GetStudentByNumber(ctx context.Context, input *usecase.GetStudentByNumberInput) (*usecase.GetStudentByNumberOutput, error) {
	return mock(ctx, input)
}

var _ usecase.StudentsGetter = GetStudents(nil)

type GetStudents func(ctx context.Context) (*usecase.GetStudentsOutput, error)

func (mock GetStudents) GetStudents(ctx context.Context) (*usecase.GetStudentsOutput, error) {
	return mock(ctx)
}

var _ usecase.StudentCreator = CreateStudent(nil)

type CreateStudent func(ctx context.Context, input *usecase.CreateStudentInput) (*usecase.CreateStudentOutput, error)

func (mock CreateStudent) CreateStudent(ctx context.Context, input *usecase.CreateStudentInput) (*usecase.CreateStudentOutput, error) {
	return mock(ctx, input)
}

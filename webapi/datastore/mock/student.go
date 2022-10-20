package mock

import (
	"context"

	"github.com/is-hoku/goa-sample/webapi/repository"
)

type mockTx string

func (tx mockTx) Commit() error {
	return nil
}

func (tx mockTx) Rollback() error {
	return nil
}

type GetStudentByNumber func(context.Context, *repository.GetStudentByNumberInput) (*repository.GetStudentByNumberOutput, error)

func (mock GetStudentByNumber) GetStudentByNumber(ctx context.Context, input *repository.GetStudentByNumberInput) (*repository.GetStudentByNumberOutput, error) {
	return mock(ctx, input)
}

type GetStudentByID func(context.Context, *repository.GetStudentByIDInput) (*repository.GetStudentByIDOutput, error)

func (mock GetStudentByID) GetStudentByID(ctx context.Context, input *repository.GetStudentByIDInput) (*repository.GetStudentByIDOutput, error) {
	return mock(ctx, input)
}

type CreateStudent func(context.Context, *repository.CreateStudentInput) (*repository.CreateStudentOutput, error)

func (mock CreateStudent) CreateStudent(ctx context.Context, input *repository.CreateStudentInput) (*repository.CreateStudentOutput, error) {
	return mock(ctx, input)
}

func (mock CreateStudent) BeginTx(ctx context.Context) (repository.Tx, error) {
	var tx mockTx
	return tx, nil
}

type GetStudents func(context.Context) (*repository.GetStudentsOutput, error)

func (mock GetStudents) GetStudents(ctx context.Context) (*repository.GetStudentsOutput, error) {
	return mock(ctx)
}

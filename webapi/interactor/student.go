package interactor

import (
	"context"
	"time"

	"github.com/is-hoku/goa-template/webapi/gen/students"
	"github.com/is-hoku/goa-template/webapi/model"
	"github.com/is-hoku/goa-template/webapi/repository"
	"github.com/is-hoku/goa-template/webapi/usecase"
)

type StudentInteractor struct {
	Repo repository.StudentRepository
}

var _ usecase.StudentUsecase = (*StudentInteractor)(nil)

func (i *StudentInteractor) GetByNum(ctx context.Context, num int64) (*model.Student, error) {
	student, err := i.Repo.Get(ctx, num)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (i *StudentInteractor) Create(ctx context.Context, student *students.StudentBody) (*model.Student, error) {
	birth, err := time.Parse(time.RFC3339, student.DateOfBirth)
	if err != nil {
		return nil, err
	}
	expiration, err := time.Parse(time.RFC3339, student.ExpirationDate)
	if err != nil {
		return nil, err
	}
	s := &model.Student{
		Name:           student.Name,
		Ruby:           student.Ruby,
		StudentNumber:  student.StudentNumber,
		DateOfBirth:    birth,
		Address:        student.Address,
		ExpirationDate: expiration,
	}
	resStudent, err := i.Repo.Set(ctx, s)
	if err != nil {
		return nil, err
	}
	return resStudent, nil
}

func (i *StudentInteractor) GetAll(ctx context.Context) ([]*model.Student, error) {
	students, err := i.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return students, nil
}

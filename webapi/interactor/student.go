package interactor

import (
	"context"

	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/is-hoku/goa-sample/webapi/repository"
	"github.com/is-hoku/goa-sample/webapi/usecase"
)

type StudentInteractor struct {
	Repo repository.StudentRepository
}

func NewStudentInteractor(repo repository.StudentRepository) *StudentInteractor {
	return &StudentInteractor{
		Repo: repo,
	}
}

var _ usecase.StudentUsecase = (*StudentInteractor)(nil)

func (i *StudentInteractor) GetByNumber(ctx context.Context, number uint32) (*model.Student, error) {
	student, err := i.Repo.GetByNumber(ctx, number)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (i *StudentInteractor) Create(ctx context.Context, s *model.Student) (*model.Student, error) {
	tx, err := i.Repo.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	studentID, err := i.Repo.Set(ctx, s)
	if err != nil {
		return nil, err
	}
	student, err := i.Repo.GetByID(ctx, studentID)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (i *StudentInteractor) GetAll(ctx context.Context) ([]*model.Student, error) {
	students, err := i.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return students, nil
}

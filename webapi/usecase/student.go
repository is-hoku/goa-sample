package usecase

import (
	"context"

	"github.com/is-hoku/goa-template/webapi/model"
)

type StudentUsecase interface {
	GetByNumber(context.Context, uint32) (*model.Student, error)
	Create(context.Context, *model.Student) (*model.Student, error)
	GetAll(context.Context) ([]*model.Student, error)
}

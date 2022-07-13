package usecase

import (
	"context"

	"github.com/is-hoku/goa-template/webapi/gen/students"
	"github.com/is-hoku/goa-template/webapi/model"
)

type StudentUsecase interface {
	GetByNum(context.Context, int64) (*model.Student, error)
	Create(context.Context, *students.StudentBody) (*model.Student, error)
	GetAll(context.Context) ([]*model.Student, error)
}

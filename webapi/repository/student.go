package repository

import (
	"context"

	"github.com/is-hoku/goa-template/webapi/model"
)

type StudentRepository interface {
	Get(context.Context, int64) (*model.Student, error)
	Set(context.Context, *model.Student) (*model.Student, error)
	GetAll(context.Context) ([]*model.Student, error)
	Close() error
}

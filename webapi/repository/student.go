package repository

import (
	"context"

	"github.com/is-hoku/goa-sample/webapi/model"
)

type StudentRepository interface {
	GetByNumber(context.Context, int32) (*model.Student, error)
	GetByID(context.Context, int64) (*model.Student, error)
	Set(context.Context, *model.Student) (int64, error)
	GetAll(context.Context) ([]*model.Student, error)
	TxBeginner
}

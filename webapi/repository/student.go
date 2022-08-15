package repository

import (
	"context"

	"github.com/is-hoku/goa-sample/webapi/model"
)

type StudentRepository interface {
	GetByNumber(context.Context, uint32) (*model.Student, error)
	GetByID(context.Context, uint64) (*model.Student, error)
	Set(context.Context, *model.Student) (uint64, error)
	GetAll(context.Context) ([]*model.Student, error)
	TxBeginner
}

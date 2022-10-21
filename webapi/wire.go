//go:build wireinject
// +build wireinject

package sample

import (
	"context"

	"github.com/google/wire"
	"github.com/is-hoku/goa-sample/webapi/datastore"
	"github.com/is-hoku/goa-sample/webapi/interactor"
)

func NewStudentApp(ctx context.Context) (*StudentApp, error) {
	wire.Build(
		newStudentApp,
		newMySQLConfig,
		datastore.StudentSet,
		interactor.StudentSet,
	)
	return &StudentApp{}, nil
}

package repository

import "context"

type Tx interface {
	Commit() error
	Rollback() error
}

type TxBeginner interface {
	BeginTx(context.Context) (Tx, error)
}

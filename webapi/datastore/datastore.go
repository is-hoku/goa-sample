package datastore

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/is-hoku/goa-sample/webapi/repository"
)

type DB struct {
	*sql.DB
}

func NewMySQL(config *mysql.Config) (*DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Printf("Could not connect to DB: %s", err)
		return nil, err
	}
	return &DB{db}, err
}

func (db *DB) BeginTx(ctx context.Context) (repository.Tx, error) {
	tx, err := db.DB.BeginTx(ctx, nil)
	return tx, err
}

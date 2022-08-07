package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/is-hoku/goa-sample/webapi/repository"
)

type DBHandler struct {
	DB *sql.DB
}

type Config struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (c Config) DNS() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
	)
}

func NewDB(config *Config) (*DBHandler, error) {
	db, err := sql.Open("mysql", config.DNS())
	if err != nil {
		log.Printf("Could not connect to DB: %s", err)
		return nil, err
	}
	return &DBHandler{db}, err
}

func (db *DBHandler) Close() error {
	return db.DB.Close()
}

func (db *DBHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
	tx, err := db.DB.BeginTx(ctx, nil)
	return tx, err
}

package testutil

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/is-hoku/goa-sample/webapi/datastore"
)

func NewTestDBHandler(ctx context.Context) (*datastore.DBHandler, error) {
	config1 := &datastore.Config{
		User:     os.Getenv("TEST_DB_USER"),
		Password: os.Getenv("TEST_DB_PASS"),
		Host:     os.Getenv("TEST_DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
	// 既存の DB に接続
	handler1, err := datastore.NewDB(config1)
	if err != nil {
		log.Fatalf("Could not generate db: %s", err)
	}
	defer handler1.DB.Close()
	if _, err := handler1.DB.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", os.Getenv("TEST_DB_NAME"))); err != nil {
		log.Fatalf("Could not create test db: %s\n", err)
		return nil, err
	}
	config2 := &datastore.Config{
		User:     os.Getenv("TEST_DB_USER"),
		Password: os.Getenv("TEST_DB_PASS"),
		Host:     os.Getenv("TEST_DB_HOST"),
		Port:     os.Getenv("TEST_DB_PORT"),
		DBName:   os.Getenv("TEST_DB_NAME"),
	}
	// テスト用 DB
	handler, err := datastore.NewDB(config2)
	if err != nil {
		log.Fatalf("Could not generate db: %s", err)
	}
	return handler, nil
}

func CreateTestStudentTable(ctx context.Context, handler *datastore.DBHandler) error {
	createTableQuery := "CREATE TABLE IF NOT EXISTS `students` (" +
		"`id` BIGINT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT," +
		"`name` VARCHAR(128) NOT NULL," +
		"`ruby` VARCHAR(128) NOT NULL," +
		"`student_number` INT UNSIGNED NOT NULL UNIQUE," +
		"`date_of_birth` DATETIME NOT NULL," +
		"`address` VARCHAR(256) NOT NULL," +
		"`expiration_date` DATETIME NOT NULL," +
		"`created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"`updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP);"
	if _, err := handler.DB.ExecContext(ctx, createTableQuery); err != nil {
		log.Fatalf("Could not create test table: %s\n", err)
		return err
	}
	return nil
}

func DeleteTestDB(ctx context.Context, handler *datastore.DBHandler) error {
	dbName := os.Getenv("TEST_DB_NAME")
	if _, err := handler.DB.ExecContext(ctx, fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", dbName)); err != nil {
		log.Fatalf("Could not delete test db: %s\n", err)
		return err
	}
	return nil
}

func TruncateAll(ctx context.Context, handler *datastore.DBHandler) error {
	rows, err := handler.DB.QueryContext(ctx, "SHOW TABLES")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return err
		}
		if _, err := handler.DB.ExecContext(ctx, fmt.Sprintf("TRUNCATE `%s`", tableName)); err != nil {
			return err
		}
	}
	return rows.Err()
}

func InsertTestStudentsData(ctx context.Context, handler *datastore.DBHandler, students []*datastore.SetStudentParams) error {
	tx, err := handler.DB.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("Could not begin a transaction: %s\n", err)
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.PrepareContext(ctx, "INSERT INTO `students` (`name`, `ruby`, `student_number`, `date_of_birth`, `address`, `expiration_date`) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatalf("Could not create a prepared statement: %s\n", err)
		return err
	}
	defer stmt.Close()
	for _, s := range students {
		_, err := stmt.ExecContext(ctx, s.Name, s.Ruby, s.StudentNumber, s.DateOfBirth, s.Address, s.ExpirationDate)
		if err != nil {
			log.Fatalf("Could not execute prepared statement: %s\n", err)
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatalf("Could not commit the transaction: %s\n", err)
		return err
	}
	return nil
}

package datastore

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

func newTestMySQLConfig() (*mysql.Config, error) {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.User = os.Getenv("TEST_DB_USER")
	config.Passwd = os.Getenv("TEST_DB_PASS")
	config.Addr = net.JoinHostPort(os.Getenv("TEST_DB_HOST"), os.Getenv("TEST_DB_PORT"))
	config.DBName = os.Getenv("TEST_DB_NAME")
	config.Timeout = 30 * time.Second
	config.RejectReadOnly = true
	config.ParseTime = true
	return config, nil
}

func newTestDB(ctx context.Context) (*DB, error) {
	config1 := mysql.NewConfig()
	config1.Net = "tcp"
	config1.User = os.Getenv("TEST_DB_USER")
	config1.Passwd = os.Getenv("TEST_DB_PASS")
	config1.Addr = net.JoinHostPort(os.Getenv("TEST_DB_HOST"), os.Getenv("DB_PORT"))
	config1.DBName = os.Getenv("DB_NAME")
	config1.Timeout = 30 * time.Second
	config1.RejectReadOnly = true
	config1.ParseTime = true
	// 既存の DB に接続
	testDB1, err := NewMySQL(config1)
	if err != nil {
		return nil, err
	}
	defer testDB1.Close()
	if _, err := testDB1.DB.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", os.Getenv("TEST_DB_NAME"))); err != nil {
		return nil, err
	}
	config2, err := newTestMySQLConfig()
	if err != nil {
		return nil, err
	}
	testDB2, err := NewMySQL(config2)
	if err != nil {
		return nil, err
	}
	return testDB2, nil
}

func createTestStudentsTable(ctx context.Context, db *DB) error {
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
	if _, err := db.ExecContext(ctx, createTableQuery); err != nil {
		log.Fatalf("Could not create test table: %s\n", err)
		return err
	}
	return nil
}

func deleteTestDB(ctx context.Context, db *DB) error {
	dbName := os.Getenv("TEST_DB_NAME")
	if _, err := db.ExecContext(ctx, fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", dbName)); err != nil {
		log.Fatalf("Could not delete test db: %s\n", err)
		return err
	}
	return nil
}

func truncateAll(ctx context.Context, db *DB) error {
	rows, err := db.QueryContext(ctx, "SHOW TABLES")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return err
		}
		if _, err := db.ExecContext(ctx, fmt.Sprintf("TRUNCATE `%s`", tableName)); err != nil {
			return err
		}
	}
	return rows.Err()
}

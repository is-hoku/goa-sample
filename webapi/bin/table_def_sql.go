package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/is-hoku/goa-sample/webapi/datastore"
)

func main() {
	config := &datastore.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
	db, err := datastore.NewDB(config)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	rows, err := db.DB.QueryContext(ctx, "SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	var tables []string
	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			log.Fatal(err)
		}
		if tableName != "users" { // users は DB が自動生成したテーブル
			tables = append(tables, tableName)
		}
	}
	for _, table := range tables {
		row := db.DB.QueryRowContext(ctx, "SHOW CREATE TABLE `"+table+"`")
		var tableName, schema string
		err = row.Scan(&tableName, &schema)
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create("./datastore/table/" + table + ".sql")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = file.WriteString(schema)
		if err != nil {
			log.Fatal(err)
		}
	}
}

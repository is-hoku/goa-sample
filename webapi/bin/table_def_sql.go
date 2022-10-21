package main

import (
	"context"
	"log"
	"net"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/is-hoku/goa-sample/webapi/datastore"
)

func newMySQLConfig() *mysql.Config {
	config := mysql.NewConfig()
	config.Net = "tcp"
	config.User = os.Getenv("DB_USER")
	config.Passwd = os.Getenv("DB_PASS")
	config.Addr = net.JoinHostPort(os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	config.DBName = os.Getenv("DB_NAME")
	config.Timeout = 30 * time.Second
	config.RejectReadOnly = true
	config.ParseTime = true
	return config
}

func main() {
	config := newMySQLConfig()
	db, err := datastore.NewMySQL(config)
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

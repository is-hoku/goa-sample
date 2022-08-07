package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/is-hoku/goa-sample/webapi/datastore"
	"github.com/is-hoku/goa-sample/webapi/model"
	"github.com/joho/godotenv"
	ddlmaker "github.com/kayac/ddl-maker"
	"github.com/schemalex/schemalex/diff"
)

func main() {
	filePath := "migration.sql"
	conf := ddlmaker.Config{
		DB: ddlmaker.DBConfig{
			Driver:  "mysql",
			Engine:  "InnoDB",
			Charset: "utf8mb4",
		},
		OutFilePath: filePath,
	}

	dm, err := ddlmaker.New(conf)
	if err != nil {
		log.Println(err.Error())
		return
	}

	structs := []interface{}{
		model.Student{},
	}

	dm.AddStruct(structs...)

	err = dm.Generate()
	if err != nil {
		log.Println(err.Error())
		return
	}

	ddl, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// あとでかきなおす
	err = godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Could not load .env: %s", err)
	}
	config := &datastore.Config{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
	}
	handler, err := datastore.New(config)
	handler.Description()

	var buf bytes.Buffer
	diff.Strings(&buf, ddl)
}

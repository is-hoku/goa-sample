package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/is-hoku/goa-sample/webapi/model"
	ddlmaker "github.com/kayac/ddl-maker"
)

func main() {
	var (
		driver      string
		engine      string
		charset     string
		outFilePath string
	)
	currentTime := time.Now().UTC().Format("2006_01_02_150405")
	flag.StringVar(&driver, "d", "", "set driver")
	flag.StringVar(&driver, "driver", "", "set driver")
	flag.StringVar(&outFilePath, "o", fmt.Sprintf("../mysql/migration/%s.sql", currentTime), "set ddl output file path")
	flag.StringVar(&outFilePath, "outfile", fmt.Sprintf("../mysql/migration/%s.sql", currentTime), "set ddl output file path")
	flag.StringVar(&engine, "e", "InnoDB", "set driver engine")
	flag.StringVar(&engine, "engine", "InnoDB", "set driver engine")
	flag.StringVar(&charset, "c", "utf8mb4", "set driver charset")
	flag.StringVar(&charset, "charset", "utf8mb4", "set driver charset")
	flag.Parse()

	if driver == "" {
		log.Println("Please set driver name. -d or -driver")
		return
	}
	if outFilePath == "" {
		log.Println("Please set outFilePath. -o or -outfile")
		return
	}

	conf := ddlmaker.Config{
		DB: ddlmaker.DBConfig{
			Driver:  driver,
			Engine:  engine,
			Charset: charset,
		},
		OutFilePath: outFilePath,
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
}
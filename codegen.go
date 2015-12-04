package main

import (
	"bytes"
	"database/sql"
	"flag"
	"fmt"
	"gopkg.in/mgutz/dat.v1"
	"gopkg.in/mgutz/dat.v1/sqlx-runner"
	"log"
	"os"
	"time"
)

type DBSchema struct {
	Column   string        `db:"column_name"`
	UpColumn string        `db:"up_column_name"`
	DataType string        `db:"data_type"`
	MaxLen   dat.NullInt64 `db:"character_maximum_length"`
}

var (
	DB         *runner.DB
	dirname    string
	dbname     string
	sqltable   string
	sqlas      string
	tablename  string
	Tablename  string
	html       bool
	form       bool
	controller bool
	backend    bool
	schema     []*DBSchema
)

type Today struct {
	Today dat.NullTime `db:"now"`
}

func _initDB() {

	db, err := sql.Open("postgres", Config.DataSourceName)
	if err != nil {
		fmt.Println("SQL Open", err.Error())
	}
	runner.MustPing(db)

	// set to reasonable values for production
	db.SetMaxIdleConns(4)
	db.SetMaxOpenConns(16)

	// set this to enable interpolation
	dat.EnableInterpolation = true

	// set to check things like sessions closing.
	// Should be disabled in production/release builds.
	dat.Strict = false

	// Log any query over 10ms as warnings. (optional)
	runner.LogQueriesThreshold = 10 * time.Millisecond

	DB = runner.NewDB(db, "postgres")

	// DoO a test run against the DB
	var _res Today
	dbErr := DB.SQL("select now()").QueryStruct(&_res)
	if dbErr != nil {
		log.Fatalln(dbErr.Error())
	}

	err = DB.SQL(`
	select column_name, data_type, character_maximum_length
	from INFORMATION_SCHEMA.COLUMNS 
	where table_name = $1`, sqltable).QueryStructs(&schema)

	if err != nil {
		fmt.Println("SQL Schema Query:", err.Error())
	}

	for _, row := range schema {
		row.UpColumn = UpperFirst(row.Column)
	}
}

func generateHTML() {

	generateHTML_list(fmt.Sprintf("%s/%s.list.html", dirname, tablename))
	generateHTML_edit(fmt.Sprintf("%s/%s.edit.html", dirname, tablename))
	generateHTML_new(fmt.Sprintf("%s/%s.new.html", dirname, tablename))
}

func UpperFirst(s string) string {
	byt := []byte(s)
	firstChar := bytes.ToUpper([]byte{byt[0]})
	rest := byt[1:]
	return string(bytes.Join([][]byte{firstChar, rest}, nil))
}

func main() {

	_loadConfig()

	flag.StringVar(&dirname, "out", "generated", "(Optional) Directory to place generated code into")
	flag.StringVar(&sqltable, "t", "", "Name of the SQL table to use")
	flag.StringVar(&sqlas, "as", "", "(Optional) name of the table Object   (default = same as SQL table name)")
	flag.BoolVar(&html, "html", false, "Generate HTML ?")
	flag.Parse()

	if dirname == "" {
		log.Fatalln("Must define an output directory to place generated code into")
	}

	if sqltable == "" {
		log.Fatalln("No table defined")
	} else {
		tablename = sqltable
		if sqlas != "" {
			tablename = sqlas
		}
		Tablename = UpperFirst(tablename)
	}

	_initDB()

	// Create the generation dir if not already there, ignore errors
	os.Mkdir(dirname, os.ModePerm)

	if html {
		generateHTML()
	}

}

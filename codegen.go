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
	"regexp"
	"time"
)

type DBSchema struct {
	Column   string        `db:"column_name"`
	UpColumn string        `db:"up_column_name"`
	DataType string        `db:"data_type"`
	MaxLen   dat.NullInt64 `db:"character_maximum_length"`
}

var (
	DB           *runner.DB
	dirname      string
	dbname       string
	sqltable     string
	sqlas        string
	tablename    string
	Tablename    string
	html         bool
	form         bool
	gotype       bool
	gorest       bool
	schema       []*DBSchema
	column_list  []string
	column_names string
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

	column_names = ""
	for i, row := range schema {
		if row.Column == "id" {
			row.UpColumn = "ID"
		} else {
			row.UpColumn = UpperFirst(row.Column)
		}
		column_list = append(column_list, row.Column)
		if i > 0 {
			column_names += ","
		}
		column_names += fmt.Sprintf("\"%s\"", row.Column)
	}
}

func generateHTML() {

	generateHTML_list(fmt.Sprintf("%s/%s.list.html", dirname, tablename))
	generateHTML_edit(fmt.Sprintf("%s/%s.edit.html", dirname, tablename))
	generateHTML_new(fmt.Sprintf("%s/%s.new.html", dirname, tablename))
}

func generateForm() {
	generate_Formly(fmt.Sprintf("%s/%s.form.js", dirname, tablename))
}

func generateGoType() {

	fmt.Println(generate_GoType())
}

func generateGoRest() {

	generate_Go_REST(fmt.Sprintf("%s/rest_%s.go", dirname, tablename))
}

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

func CamelCase(src string) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		if idx > 0 {
			chunks[idx] = bytes.Title(val)
		}
	}
	return string(bytes.Join(chunks, nil))
}

func UpperFirst(s string) string {
	byt := []byte(CamelCase(s))
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
	flag.BoolVar(&form, "formly", false, "Generate ngFormly Defintiions ?")
	flag.BoolVar(&gotype, "gotype", false, "Generate Go type declaration")
	flag.BoolVar(&gorest, "gorest", false, "Generate Go REST handlers for this table")
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

	if form {
		generateForm()
	}

	if gotype {
		generateGoType()
	}

	if gorest {
		generateGoRest()
	}

}

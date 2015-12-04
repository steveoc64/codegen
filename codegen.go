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

	db, err := sql.Open("postgres", fmt.Sprintf("user=postgres password=unx911zxx dbname=%s sslmode=disable", dbname))
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

	DB.SQL(`
	select column_name, data_type, character_maximum_length
	from INFORMATION_SCHEMA.COLUMNS 
	where table_name = $1`, tablename).QueryStructs(&schema)

	for _, row := range schema {
		row.UpColumn = UpperFirst(row.Column)
	}
}

func generateHTML() {

	listFile := fmt.Sprintf("%s/%ss.html", dirname, tablename)
	generateHTML_list(listFile)
}

func UpperFirst(s string) string {
	byt := []byte(s)
	firstChar := bytes.ToUpper([]byte{byt[0]})
	rest := byt[1:]
	return string(bytes.Join([][]byte{firstChar, rest}, nil))
}

func main() {

	flag.StringVar(&dirname, "out", "generated", "Directory to place generated code into")
	flag.StringVar(&dbname, "d", "cmms", "Name of the database to connect to")
	flag.StringVar(&tablename, "t", "", "Name of the table to use")
	flag.BoolVar(&html, "html", false, "Generate HTML ?")
	flag.Parse()
	if dirname == "" {
		log.Fatalln("Must define an output directory to place generated code into")
	}
	if dbname == "" {
		log.Fatalln("No database defined")
	}
	if tablename == "" {
		log.Fatalln("No table defined")
	} else {
		Tablename = UpperFirst(tablename)
	}

	_initDB()

	// Create the generation dir if not already there, ignore errors
	os.Mkdir(dirname, os.ModePerm)

	if html {
		generateHTML()
	}

}
package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	driver   = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" //same as in make file
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(driver, dbSource)
	if err != nil {
		log.Fatal("can not connect to database")
	}
	testQueries = New(conn)

	os.Exit(m.Run())
}

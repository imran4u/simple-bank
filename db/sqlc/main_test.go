package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/imran4u/simple-bank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("Can't load config", err)
	}

	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to database")
	}
	testQueries = New(testDb)

	os.Exit(m.Run())
}

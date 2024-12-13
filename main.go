package main

import (
	"database/sql"
	"log"

	"github.com/imran4u/simple-bank/api"
	db "github.com/imran4u/simple-bank/db/sqlc"

	_ "github.com/lib/pq" // this lib is important to connect with postgres db
)

const (
	driver   = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" //same as in make file
	address  = "0.0.0.0:8080"
)

func main() {

	conn, err := sql.Open(driver, dbSource)
	if err != nil {
		log.Fatal("can not connect to database")
	}
	store := db.Store{
		Queries: db.New(conn),
		Db:      conn}
	server := api.NewServer(&store)

	err = server.Start(address)
	if err != nil {
		log.Fatal("Fail to start server : ", err.Error())

	}
}

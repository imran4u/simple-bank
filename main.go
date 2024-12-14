package main

import (
	"database/sql"
	"log"

	"github.com/imran4u/simple-bank/api"
	db "github.com/imran4u/simple-bank/db/sqlc"
	"github.com/imran4u/simple-bank/util"

	_ "github.com/lib/pq" // this lib is important to connect with postgres db
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Can't load config file", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("can not connect to database")
	}
	store := db.SQLStore{
		Queries: db.New(conn),
		Db:      conn}
	server := api.NewServer(&store)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("Fail to start server : ", err.Error())

	}
}

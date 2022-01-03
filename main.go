package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/matiolsz/simplebank/api"
	db "github.com/matiolsz/simplebank/db/sqlc"
	"github.com/matiolsz/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

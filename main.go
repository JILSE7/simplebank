package main

import (
	"database/sql"
	"log"

	"github.com/JILSE7/simplebank/api"
	db "github.com/JILSE7/simplebank/db/sqlc"
	"github.com/JILSE7/simplebank/utils"
	_ "github.com/lib/pq" // pq to talk with the DB important
)

func main() {
	config, err := utils.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	connection, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot conncet to db", err)
	}

	store := db.NewStore(connection)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("can´t start the server", err)
	}

}

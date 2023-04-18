package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/keidarcy/simple-bank/api"
	db "github.com/keidarcy/simple-bank/db/sqlc"
	"github.com/keidarcy/simple-bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	fmt.Printf("%v", config)

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

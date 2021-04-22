package main

import (
	"database/sql"
	"log"

	"github.com/dnswd/kanal/api"
	db "github.com/dnswd/kanal/db/sqlc"
	"github.com/dnswd/kanal/util"
	_ "github.com/lib/pq"
)


func main() {
	config, err := util.ReadConfig(".")
	if err != nil {
		log.Fatal("Error while reading configs", err)
	}

	conn, err := sql.Open(config.SourceType, config.Source)
	if err != nil {
		log.Fatal("Error while connecting to DB", err)
	}

	store := db.NewStore(conn)
	driver := api.InitDriver(config, store)
	err = driver.Listen()
	if err != nil {
		log.Fatal("Error while starting the server", err)
	}
}

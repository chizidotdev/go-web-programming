package data

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "dbname=go-chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"subscription-service/internal/database"
)

func init() {
	db := database.InitDB()
	_ = db.Ping()
}

func main() {
	// create sessions

	// create channels

	// create waitgroup

	// set up the application config

	// set up mail

	// listen for web connections
}

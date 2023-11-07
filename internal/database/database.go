package database

import (
	"database/sql"
	"log"
	"os"
	"time"
)

// Create connects to postgres, and backs off until a connection
func Create() *sql.DB {
	conn := connect()
	if conn == nil {
		log.Panic("can't connect to database")
	}
	return conn
}

// connect connects to postgres, and backs off until a connection
func connect() *sql.DB {
	counts := 0

	dsn := os.Getenv("DSN")

	for {
		connection, err := open(dsn)
		if err != nil {
			log.Println("postgres not yet ready...")
		} else {
			log.Print("connected to database!")
			return connection
		}

		if counts > 10 {
			return nil
		}

		log.Print("Backing off for 1 second")
		time.Sleep(1 * time.Second)
		counts++

		continue
	}
}

// open opens a connection to postgres
func open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

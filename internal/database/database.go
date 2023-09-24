package database

import (
	"database/sql"
	"log"
	"os"
	"time"
)

// InitDB gets a connection to the database.
func InitDB() *sql.DB {
	conn := connectDB()
	if conn == nil {
		log.Panic("can't connect to the database")
	}
	return conn
}

func connectDB() *sql.DB {
	counts := 0
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err == nil {
			log.Print("connected to the database!")
			return connection
		}

		log.Print("postgres not ready yet!")

		if counts > 10 {
			return nil
		}

		log.Print("backing off for 1 second")
		time.Sleep(1 * time.Second)
		counts++

		continue
	}
}

func openDB(dsn string) (*sql.DB, error) {
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

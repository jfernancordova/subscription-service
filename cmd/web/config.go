package main

import (
	"database/sql"
	"log"
	"subscription-service/data"
	"sync"

	"github.com/alexedwards/scs/v2"
)

// Config holds application configuration
type Config struct {
	Session       *scs.SessionManager
	DB            *sql.DB
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	Wait          *sync.WaitGroup
	Models        data.Models
	Mailer        Mail
	ErrorChan     chan error
	ErrorChanDone chan bool
}

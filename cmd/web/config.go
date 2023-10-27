package main

import (
	"database/sql"
	"log"
	"subscription-service/data"
	"sync"

	"github.com/alexedwards/scs/v2"
)

// config holds application configuration
type config struct {
	session       *scs.SessionManager
	db            *sql.DB
	infoLog       *log.Logger
	errorLog      *log.Logger
	wait          *sync.WaitGroup
	models        data.Models
	mailer        Mail
	errorChan     chan error
	errorChanDone chan bool
}

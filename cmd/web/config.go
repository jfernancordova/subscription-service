package main

import (
	"database/sql"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type config struct {
	session  *scs.SessionManager
	db       *sql.DB
	infoLog  *log.Logger
	errorLog *log.Logger
	wait     *sync.WaitGroup
}

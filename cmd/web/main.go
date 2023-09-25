package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/alexedwards/scs/v2"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"

	"subscription-service/internal/database"
	s "subscription-service/internal/session"
)

var db *sql.DB
var session *scs.SessionManager

func init() {
	db = database.Init()
	session = s.Init()
}

func main() {
	wg := sync.WaitGroup{}
	app := config{
		session:  session,
		db:       db,
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		wait:     &wg,
	}

	// set up mail

	// listen for web connections
	app.serve()
}

func (app *config) serve() {
	// start http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", "80"),
		Handler: app.routes(),
	}
	app.infoLog.Println("Starting web server...")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

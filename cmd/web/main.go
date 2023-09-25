package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

	go app.listenShutdown()

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

func (app *config) listenShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *config) shutdown() {
	// perform any cleanup task
	app.infoLog.Println("run cleanup task...")

	// block until waitgroup is empty
	app.wait.Wait()

	app.infoLog.Println("closing channels and shutting down application...")
}
